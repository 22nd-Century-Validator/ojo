package e2e

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	oracletypes "github.com/ojo-network/ojo/x/oracle/types"
)

var (
	oracleAcceptList = []oracletypes.Denom{
		{BaseDenom: "uumee", SymbolDenom: "UMEE", Exponent: 6},
		{BaseDenom: "ibc/1", SymbolDenom: "ATOM", Exponent: 6},
		{BaseDenom: "ibc/2", SymbolDenom: "USDC", Exponent: 6},
		{BaseDenom: "ibc/3", SymbolDenom: "DAI", Exponent: 18},
		{BaseDenom: "ibc/4", SymbolDenom: "ETH", Exponent: 18},
		{BaseDenom: "ibc/5", SymbolDenom: "WBTC", Exponent: 8},
		{BaseDenom: "ibc/6", SymbolDenom: "BNB", Exponent: 18},
		{BaseDenom: "ibc/7", SymbolDenom: "stATOM", Exponent: 6},
		{BaseDenom: "ibc/8", SymbolDenom: "stOSMO", Exponent: 6},
		{BaseDenom: "ibc/9", SymbolDenom: "OSMO", Exponent: 6},
		{BaseDenom: "ibc/9", SymbolDenom: "IST", Exponent: 6},
	}

	oracleRewardBands = []oracletypes.RewardBand{
		{SymbolDenom: "UMEE", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "ATOM", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "USDC", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "DAI", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "ETH", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "WBTC", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "BNB", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "stATOM", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "stOSMO", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "OSMO", RewardBand: sdk.MustNewDecFromStr("1.0")},
		{SymbolDenom: "IST", RewardBand: sdk.MustNewDecFromStr("1.0")},
	}
)
