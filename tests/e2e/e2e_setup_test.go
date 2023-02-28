package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	tmconfig "github.com/tendermint/tendermint/config"
	tmjson "github.com/tendermint/tendermint/libs/json"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	appparams "github.com/ojo-network/ojo/app/params"
	"github.com/ojo-network/ojo/client"
	oracletypes "github.com/ojo-network/ojo/x/oracle/types"
)

const (
	ojoContainerRepo = "ojo"
	ojoP2pPort       = "26656"
	ojoTmrpcPort     = "26657"
	ojoGrpcPort      = "9090"

	priceFeederContainerRepo = "ghcr.io/umee-network/price-feeder-e2e"
	priceFeederServerPort    = "7171/tcp"

	initBalanceStr = "510000000000" + appparams.BondDenom
)

type IntegrationTestSuite struct {
	suite.Suite

	tmpDirs             []string
	chain               *chain
	dkrPool             *dockertest.Pool
	dkrNet              *dockertest.Network
	priceFeederResource *dockertest.Resource
	valResources        []*dockertest.Resource
	ojoClient           *client.OjoClient
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up e2e integration test suite...")
	appparams.SetAddressPrefixes()

	var err error
	s.chain, err = newChain()
	s.Require().NoError(err)

	s.T().Logf("starting e2e infrastructure; chain-id: %s; datadir: %s", s.chain.id, s.chain.dataDir)

	s.dkrPool, err = dockertest.NewPool("")
	s.Require().NoError(err)

	s.dkrNet, err = s.dkrPool.CreateNetwork(fmt.Sprintf("%s-testnet", s.chain.id))
	s.Require().NoError(err)

	s.initNodes()
	s.initGenesis()
	s.initValidatorConfigs()
	s.runValidators()
	s.initOjoClient()
	s.delegatePriceFeederVoting()
	s.runPriceFeeder()
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e integration test suite...")

	s.Require().NoError(s.dkrPool.Purge(s.priceFeederResource))

	for _, vc := range s.valResources {
		s.Require().NoError(s.dkrPool.Purge(vc))
	}

	s.Require().NoError(s.dkrPool.RemoveNetwork(s.dkrNet))

	os.RemoveAll(s.chain.dataDir)
	for _, td := range s.tmpDirs {
		os.RemoveAll(td)
	}
}

func (s *IntegrationTestSuite) initNodes() {
	s.Require().NoError(s.chain.createAndInitValidators(3))

	// initialize a genesis file for the first validator
	val0ConfigDir := s.chain.validators[0].configDir()
	for _, val := range s.chain.validators {
		valAddr, err := val.keyInfo.GetAddress()
		s.Require().NoError(err)
		s.Require().NoError(
			addGenesisAccount(val0ConfigDir, "", initBalanceStr, valAddr),
		)
	}

	// copy the genesis file to the remaining validators
	for _, val := range s.chain.validators[1:] {
		_, err := copyFile(
			filepath.Join(val0ConfigDir, "config", "genesis.json"),
			filepath.Join(val.configDir(), "config", "genesis.json"),
		)
		s.Require().NoError(err)
	}
}

func (s *IntegrationTestSuite) initGenesis() {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(s.chain.validators[0].configDir())
	config.Moniker = s.chain.validators[0].moniker

	genFilePath := config.GenesisFile()
	s.T().Log("starting e2e infrastructure; validator_0 config:", genFilePath)
	appGenState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFilePath)
	s.Require().NoError(err)

	// Oracle
	var oracleGenState oracletypes.GenesisState
	s.Require().NoError(cdc.UnmarshalJSON(appGenState[oracletypes.ModuleName], &oracleGenState))

	oracleGenState.Params.HistoricStampPeriod = 5
	oracleGenState.Params.MaximumPriceStamps = 4
	oracleGenState.Params.MedianStampPeriod = 20
	oracleGenState.Params.MaximumMedianStamps = 2
	oracleGenState.Params.AcceptList = oracleAcceptList
	oracleGenState.Params.RewardBands = oracleRewardBands

	bz, err := cdc.MarshalJSON(&oracleGenState)
	s.Require().NoError(err)
	appGenState[oracletypes.ModuleName] = bz

	// Gov
	var govGenState govtypesv1.GenesisState
	s.Require().NoError(cdc.UnmarshalJSON(appGenState[govtypes.ModuleName], &govGenState))

	var votingPeroid = 5 * time.Second
	govGenState.VotingParams.VotingPeriod = &votingPeroid

	bz, err = cdc.MarshalJSON(&govGenState)
	s.Require().NoError(err)
	appGenState[govtypes.ModuleName] = bz

	// Genesis Txs
	var genUtilGenState genutiltypes.GenesisState
	s.Require().NoError(cdc.UnmarshalJSON(appGenState[genutiltypes.ModuleName], &genUtilGenState))

	genTxs := make([]json.RawMessage, len(s.chain.validators))
	for i, val := range s.chain.validators {
		var createValmsg sdk.Msg
		if i == 2 {
			createValmsg, err = val.buildCreateValidatorMsg(majorityValidatorStake)
		} else {
			createValmsg, err = val.buildCreateValidatorMsg(minorityValidatorStake)
		}
		s.Require().NoError(err)

		signedTx, err := val.signMsg(createValmsg)
		s.Require().NoError(err)

		txRaw, err := cdc.MarshalJSON(signedTx)
		s.Require().NoError(err)

		genTxs[i] = txRaw
	}

	genUtilGenState.GenTxs = genTxs

	bz, err = cdc.MarshalJSON(&genUtilGenState)
	s.Require().NoError(err)
	appGenState[genutiltypes.ModuleName] = bz

	bz, err = json.MarshalIndent(appGenState, "", "  ")
	s.Require().NoError(err)

	genDoc.AppState = bz

	bz, err = tmjson.MarshalIndent(genDoc, "", "  ")
	s.Require().NoError(err)

	// write the updated genesis file to each validator
	for _, val := range s.chain.validators {
		writeFile(filepath.Join(val.configDir(), "config", "genesis.json"), bz)
	}
}

func (s *IntegrationTestSuite) initValidatorConfigs() {
	for i, val := range s.chain.validators {
		tmCfgPath := filepath.Join(val.configDir(), "config", "config.toml")

		vpr := viper.New()
		vpr.SetConfigFile(tmCfgPath)
		s.Require().NoError(vpr.ReadInConfig())

		valConfig := tmconfig.DefaultConfig()
		s.Require().NoError(vpr.Unmarshal(valConfig))

		valConfig.P2P.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%s", ojoP2pPort)
		valConfig.P2P.AddrBookStrict = false
		valConfig.P2P.ExternalAddress = fmt.Sprintf("%s:%s", val.instanceName(), ojoP2pPort)
		valConfig.RPC.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%s", ojoTmrpcPort)
		valConfig.StateSync.Enable = false
		valConfig.LogLevel = "info"

		var peers []string

		for j := 0; j < len(s.chain.validators); j++ {
			if i == j {
				continue
			}

			peer := s.chain.validators[j]
			peerID := fmt.Sprintf("%s@%s%d:26656", peer.nodeKey.ID(), peer.moniker, j)
			peers = append(peers, peerID)
		}

		valConfig.P2P.PersistentPeers = strings.Join(peers, ",")

		tmconfig.WriteConfigFile(tmCfgPath, valConfig)

		// set application configuration
		appCfgPath := filepath.Join(val.configDir(), "config", "app.toml")

		appConfig := srvconfig.DefaultConfig()
		appConfig.API.Enable = true
		appConfig.MinGasPrices = minGasPrice

		srvconfig.WriteConfigFile(appCfgPath, appConfig)
	}
}

func (s *IntegrationTestSuite) runValidators() {
	s.T().Log("starting ojo validator containers...")

	s.valResources = make([]*dockertest.Resource, len(s.chain.validators))
	for i, val := range s.chain.validators {
		runOpts := &dockertest.RunOptions{
			Name:      val.instanceName(),
			NetworkID: s.dkrNet.Network.ID,
			Mounts: []string{
				fmt.Sprintf("%s/:/root/.ojo", val.configDir()),
			},
			Repository: ojoContainerRepo,
		}

		// expose the first validator
		if val.index == 0 {
			runOpts.PortBindings = map[docker.Port][]docker.PortBinding{
				"1317/tcp":  {{HostIP: "", HostPort: "1317"}},
				"9090/tcp":  {{HostIP: "", HostPort: "9090"}},
				"26656/tcp": {{HostIP: "", HostPort: "26656"}},
				"26657/tcp": {{HostIP: "", HostPort: "26657"}},
			}
		}

		resource, err := s.dkrPool.RunWithOptions(runOpts, noRestart)
		s.Require().NoError(err)

		s.valResources[i] = resource
		s.T().Logf("started ojo validator container: %s", resource.Container.ID)
	}

	rpcUrl := fmt.Sprintf("tcp://localhost:%s", ojoTmrpcPort)
	rpcClient, err := rpchttp.New(rpcUrl, "/websocket")
	s.Require().NoError(err)

	s.Require().Eventually(
		func() bool {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			status, err := rpcClient.Status(ctx)
			if err != nil {
				return false
			}

			// let the node produce a few blocks
			if status.SyncInfo.CatchingUp || status.SyncInfo.LatestBlockHeight < 3 {
				return false
			}

			return true
		},
		2*time.Minute,
		time.Second,
		"Ojo node failed to produce blocks",
	)
}

func (s *IntegrationTestSuite) delegatePriceFeederVoting() {
	delegateAddr, err := s.chain.validators[1].keyInfo.GetAddress()
	s.Require().NoError(err)
	_, err = s.ojoClient.TxClient.TxDelegateFeedConsent(delegateAddr)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) runPriceFeeder() {
	s.T().Log("starting price-feeder container...")

	votingVal := s.chain.validators[2]
	votingValAddr, err := votingVal.keyInfo.GetAddress()
	s.Require().NoError(err)

	delegateVal := s.chain.validators[1]
	delegateValAddr, err := delegateVal.keyInfo.GetAddress()
	s.Require().NoError(err)

	grpcEndpoint := fmt.Sprintf("tcp://%s:%s", delegateVal.instanceName(), ojoGrpcPort)
	tmrpcEndpoint := fmt.Sprintf("http://%s:%s", delegateVal.instanceName(), ojoTmrpcPort)

	s.priceFeederResource, err = s.dkrPool.RunWithOptions(
		&dockertest.RunOptions{
			Name:       "price-feeder",
			NetworkID:  s.dkrNet.Network.ID,
			Repository: priceFeederContainerRepo,
			Mounts: []string{
				fmt.Sprintf("%s/:/root/.ojo", delegateVal.configDir()),
			},
			PortBindings: map[docker.Port][]docker.PortBinding{
				"7171/tcp": {{HostIP: "", HostPort: "7171"}},
			},
			Env: []string{
				fmt.Sprintf("PRICE_FEEDER_PASS=%s", keyringPassphrase),
				fmt.Sprintf("ACCOUNT_ADDRESS=%s", delegateValAddr),
				fmt.Sprintf("ACCOUNT_VALIDATOR=%s", sdk.ValAddress(votingValAddr)),
				fmt.Sprintf("KEYRING_DIR=%s", "/root/.ojo"),
				fmt.Sprintf("ACCOUNT_CHAIN_ID=%s", s.chain.id),
				fmt.Sprintf("RPC_GRPC_ENDPOINT=%s", grpcEndpoint),
				fmt.Sprintf("RPC_TMRPC_ENDPOINT=%s", tmrpcEndpoint),
			},
		},
		noRestart,
	)
	s.Require().NoError(err)

	endpoint := fmt.Sprintf("http://%s/api/v1/prices", s.priceFeederResource.GetHostPort(priceFeederServerPort))
	s.Require().Eventually(
		func() bool {
			resp, err := http.Get(endpoint)
			if err != nil {
				s.T().Log("Price feeder endpoint not available", err)
				return false
			}

			defer resp.Body.Close()

			bz, err := io.ReadAll(resp.Body)
			if err != nil {
				s.T().Log("Can't get price feeder response", err)
				return false
			}

			var respBody map[string]interface{}
			if err := json.Unmarshal(bz, &respBody); err != nil {
				s.T().Log("Can't unmarshal price feed", err)
				return false
			}

			prices, ok := respBody["prices"].(map[string]interface{})
			if !ok {
				s.T().Log("price feeder: no prices")
				return false
			}

			return len(prices) > 0
		},
		time.Minute,
		time.Second,
		"price-feeder not healthy",
	)

	s.T().Logf("started price-feeder container: %s", s.priceFeederResource.Container.ID)
}

func (s *IntegrationTestSuite) initOjoClient() {
	var err error
	s.ojoClient, err = client.NewOjoClient(
		s.chain.id,
		fmt.Sprintf("tcp://localhost:%s", ojoTmrpcPort),
		fmt.Sprintf("tcp://localhost:%s", ojoGrpcPort),
		"val1",
		s.chain.validators[2].mnemonic,
	)
	s.Require().NoError(err)
}

func noRestart(config *docker.HostConfig) {
	config.RestartPolicy = docker.RestartPolicy{
		Name: "no",
	}
}
