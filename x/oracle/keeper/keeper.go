package keeper

import (
	"fmt"
	"strings"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ojo-network/ojo/util/metrics"
	"github.com/ojo-network/ojo/x/oracle/types"
)

var ten = sdk.MustNewDecFromStr("10")

// Keeper of the oracle store
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	paramSpace paramstypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	distrKeeper   types.DistributionKeeper
	StakingKeeper types.StakingKeeper

	distrName        string
	telemetryEnabled bool
	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper constructs a new keeper for oracle
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	paramspace paramstypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	distrKeeper types.DistributionKeeper,
	stakingKeeper types.StakingKeeper,
	distrName string,
	telemetryEnabled bool,
	authority string,
) Keeper {
	// ensure oracle module account is set
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	// set KeyTable if it has not already been set
	if !paramspace.HasKeyTable() {
		paramspace = paramspace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		paramSpace:       paramspace,
		accountKeeper:    accountKeeper,
		bankKeeper:       bankKeeper,
		distrKeeper:      distrKeeper,
		StakingKeeper:    stakingKeeper,
		distrName:        distrName,
		telemetryEnabled: telemetryEnabled,
		authority:        authority,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetExchangeRate gets the consensus exchange rate of USD denominated in the
// denom asset from the store.
func (k Keeper) GetExchangeRate(ctx sdk.Context, symbol string) (types.PriceStamp, error) {
	store := ctx.KVStore(k.storeKey)
	symbol = strings.ToUpper(symbol)
	b := store.Get(types.GetExchangeRateKey(symbol))
	if b == nil {
		return types.PriceStamp{}, types.ErrUnknownDenom.Wrap(symbol)
	}

	priceStamp := types.PriceStamp{}
	k.cdc.MustUnmarshal(b, &priceStamp)

	return priceStamp, nil
}

// GetExchangeRateBase gets the consensus exchange rate of an asset
// in the base denom (e.g. ATOM -> uatom)
func (k Keeper) GetExchangeRateBase(ctx sdk.Context, denom string) (sdk.Dec, error) {
	var symbol string
	var exponent uint64
	// Translate the base denom -> symbol
	params := k.GetParams(ctx)
	for _, listDenom := range params.AcceptList {
		if listDenom.BaseDenom == denom {
			symbol = listDenom.SymbolDenom
			exponent = uint64(listDenom.Exponent)
			break
		}
	}
	if len(symbol) == 0 {
		return sdk.ZeroDec(), types.ErrUnknownDenom.Wrap(denom)
	}

	priceStamp, err := k.GetExchangeRate(ctx, symbol)
	if err != nil {
		return sdk.ZeroDec(), err
	}

	powerReduction := ten.Power(exponent)
	return priceStamp.ExchangeRate.Amount.Quo(powerReduction), nil
}

// SetExchangeRate sets the consensus exchange rate of USD denominated in the
// denom asset to the store.
func (k Keeper) SetExchangeRate(ctx sdk.Context, priceStamp types.PriceStamp) {
	store := ctx.KVStore(k.storeKey)
	priceStamp.ExchangeRate.Denom = strings.ToUpper(priceStamp.ExchangeRate.Denom)
	bz := k.cdc.MustMarshal(&priceStamp)

	store.Set(types.GetExchangeRateKey(priceStamp.ExchangeRate.Denom), bz)
	go metrics.RecordExchangeRate(priceStamp.ExchangeRate.Denom, priceStamp.ExchangeRate.Amount)
}

// SetExchangeRateWithEvent sets an consensus
// exchange rate to the store with ABCI event
func (k Keeper) SetExchangeRateWithEvent(ctx sdk.Context, denom string, exchangeRate sdk.Dec) error {
	blockNum := uint64(ctx.BlockHeight())
	k.SetExchangeRate(ctx, *types.NewPriceStamp(exchangeRate, denom, blockNum))
	return ctx.EventManager().EmitTypedEvent(&types.EventSetFxRate{
		Denom: denom, Rate: exchangeRate, BlockNum: blockNum,
	})
}

// IterateExchangeRates iterates over USD rates in the store.
func (k Keeper) IterateExchangeRates(ctx sdk.Context, handler func(string, types.PriceStamp) bool) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixExchangeRate)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		key := iter.Key()
		denom := string(key[len(types.KeyPrefixExchangeRate) : len(key)-1])
		priceStamp := types.PriceStamp{}

		k.cdc.MustUnmarshal(iter.Value(), &priceStamp)
		if handler(denom, priceStamp) {
			break
		}
	}
}

func (k Keeper) HasActiveExchangeRate(ctx sdk.Context, denom string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetExchangeRateKey(denom))
}

func (k Keeper) HasActiveExchangeRates(ctx sdk.Context, denoms []string) (bool, error) {
	store := ctx.KVStore(k.storeKey)
	for _, denom := range denoms {
		found := store.Has(types.GetExchangeRateKey(denom))
		if !found {
			return false, types.ErrUnknownDenom.Wrap(denom)
		}
	}

	return true, nil
}

func (k Keeper) ClearExchangeRates(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixExchangeRate)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// GetFeederDelegation gets the account address to which the validator operator
// delegated oracle vote rights.
func (k Keeper) GetFeederDelegation(ctx sdk.Context, vAddr sdk.ValAddress) (sdk.AccAddress, error) {
	// check that the given validator exists
	if val := k.StakingKeeper.Validator(ctx, vAddr); val == nil || !val.IsBonded() {
		return nil, stakingtypes.ErrNoValidatorFound.Wrapf("validator %s is not in active set", vAddr)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetFeederDelegationKey(vAddr))
	if bz == nil {
		// no delegation, so validator itself must provide price feed
		return sdk.AccAddress(vAddr), nil
	}
	return sdk.AccAddress(bz), nil
}

// SetFeederDelegation sets the account address to which the validator operator
// delegated oracle vote rights.
func (k Keeper) SetFeederDelegation(ctx sdk.Context, operator sdk.ValAddress, delegatedFeeder sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetFeederDelegationKey(operator), delegatedFeeder.Bytes())
}

type IterateFeederDelegationHandler func(delegator sdk.ValAddress, delegate sdk.AccAddress) (stop bool)

// IterateFeederDelegations iterates over the feed delegates and performs a
// callback function.
func (k Keeper) IterateFeederDelegations(ctx sdk.Context, handler IterateFeederDelegationHandler) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixFeederDelegation)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		delegator := sdk.ValAddress(iter.Key()[2:])
		delegate := sdk.AccAddress(iter.Value())

		if handler(delegator, delegate) {
			break
		}
	}
}

// GetAggregateExchangeRatePrevote retrieves an oracle prevote from the store.
func (k Keeper) GetAggregateExchangeRatePrevote(
	ctx sdk.Context,
	voter sdk.ValAddress,
) (types.AggregateExchangeRatePrevote, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetAggregateExchangeRatePrevoteKey(voter))
	if bz == nil {
		return types.AggregateExchangeRatePrevote{}, types.ErrNoAggregatePrevote.Wrap(voter.String())
	}

	var aggregatePrevote types.AggregateExchangeRatePrevote
	k.cdc.MustUnmarshal(bz, &aggregatePrevote)

	return aggregatePrevote, nil
}

// HasAggregateExchangeRatePrevote checks if a validator has an existing prevote.
func (k Keeper) HasAggregateExchangeRatePrevote(
	ctx sdk.Context,
	voter sdk.ValAddress,
) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetAggregateExchangeRatePrevoteKey(voter))
}

// SetAggregateExchangeRatePrevote set an oracle aggregate prevote to the store.
func (k Keeper) SetAggregateExchangeRatePrevote(
	ctx sdk.Context,
	voter sdk.ValAddress,
	prevote types.AggregateExchangeRatePrevote,
) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&prevote)
	store.Set(types.GetAggregateExchangeRatePrevoteKey(voter), bz)
}

// DeleteAggregateExchangeRatePrevote deletes an oracle prevote from the store.
func (k Keeper) DeleteAggregateExchangeRatePrevote(ctx sdk.Context, voter sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAggregateExchangeRatePrevoteKey(voter))
}

// IterateAggregateExchangeRatePrevotes iterates rate over prevotes in the store
func (k Keeper) IterateAggregateExchangeRatePrevotes(
	ctx sdk.Context,
	handler func(sdk.ValAddress, types.AggregateExchangeRatePrevote) bool,
) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixAggregateExchangeRatePrevote)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		voterAddr := sdk.ValAddress(iter.Key()[2:])

		var aggregatePrevote types.AggregateExchangeRatePrevote
		k.cdc.MustUnmarshal(iter.Value(), &aggregatePrevote)

		if handler(voterAddr, aggregatePrevote) {
			break
		}
	}
}

// GetAggregateExchangeRateVote retrieves an oracle prevote from the store.
func (k Keeper) GetAggregateExchangeRateVote(
	ctx sdk.Context,
	voter sdk.ValAddress,
) (types.AggregateExchangeRateVote, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetAggregateExchangeRateVoteKey(voter))
	if bz == nil {
		return types.AggregateExchangeRateVote{}, types.ErrNoAggregateVote.Wrap(voter.String())
	}

	var aggregateVote types.AggregateExchangeRateVote
	k.cdc.MustUnmarshal(bz, &aggregateVote)

	return aggregateVote, nil
}

// SetAggregateExchangeRateVote adds an oracle aggregate prevote to the store.
func (k Keeper) SetAggregateExchangeRateVote(
	ctx sdk.Context,
	voter sdk.ValAddress,
	vote types.AggregateExchangeRateVote,
) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&vote)
	store.Set(types.GetAggregateExchangeRateVoteKey(voter), bz)
}

// DeleteAggregateExchangeRateVote deletes an oracle prevote from the store.
func (k Keeper) DeleteAggregateExchangeRateVote(ctx sdk.Context, voter sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAggregateExchangeRateVoteKey(voter))
}

type IterateExchangeRateVote = func(
	voterAddr sdk.ValAddress,
	aggregateVote types.AggregateExchangeRateVote,
) (stop bool)

// IterateAggregateExchangeRateVotes iterates rate over prevotes in the store.
func (k Keeper) IterateAggregateExchangeRateVotes(
	ctx sdk.Context,
	handler IterateExchangeRateVote,
) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixAggregateExchangeRateVote)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		voterAddr := sdk.ValAddress(iter.Key()[2:])

		var aggregateVote types.AggregateExchangeRateVote
		k.cdc.MustUnmarshal(iter.Value(), &aggregateVote)

		if handler(voterAddr, aggregateVote) {
			break
		}
	}
}

// ValidateFeeder returns error if the given feeder is not allowed to feed the message.
func (k Keeper) ValidateFeeder(ctx sdk.Context, feederAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	delegate, err := k.GetFeederDelegation(ctx, valAddr)
	if err != nil {
		return err
	}
	if !delegate.Equals(feederAddr) {
		return types.ErrNoVotingPermission.Wrap(feederAddr.String())
	}

	return nil
}

func (k Keeper) IterateExchangeRatesForDenoms(ctx sdk.Context, denoms []string) (types.PriceStamps, error) {
	prices := types.PriceStamps{}
	store := ctx.KVStore(k.storeKey)
	for _, symbol := range denoms {
		symbol = strings.ToUpper(symbol)
		b := store.Get(types.GetExchangeRateKey(symbol))
		if b == nil {
			return nil, types.ErrUnknownDenom.Wrap(symbol)
		}

		priceStamp := types.PriceStamp{}
		k.cdc.MustUnmarshal(b, &priceStamp)
		prices = append(prices, priceStamp)
	}

	return prices, nil
}
