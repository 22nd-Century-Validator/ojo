package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	oracleTypes "github.com/ojo-network/ojo/x/oracle/types"
)

type OracleKeeper interface {
	IterateExchangeRatesForDenoms(ctx sdk.Context, denoms []string) (oracleTypes.PriceStamps, error)
	IterateHistoricPricesForDenoms(ctx sdk.Context, prefix []byte, denoms []string, numStamps uint) oracleTypes.PriceStamps
	HasActiveExchangeRates(ctx sdk.Context, denom []string) (bool, error)
	HasActiveHistoricalRates(ctx sdk.Context, denom []string) (bool, error)
	MaximumMedianStamps(ctx sdk.Context) (res uint64)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
