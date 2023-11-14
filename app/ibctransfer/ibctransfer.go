package ibctransfer

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	types "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	gmptypes "github.com/ojo-network/ojo/x/gmp/types"
)

type Keeper struct {
	ibctransferkeeper.Keeper

	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	paramSpace paramtypes.Subspace

	ics4Wrapper   porttypes.ICS4Wrapper
	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
	authKeeper    types.AccountKeeper
	bankKeeper    types.BankKeeper
	scopedKeeper  exported.ScopedKeeper

	gmpKeeper GmpKeeper
}

// Transfer defines a wrapper function for the ICS20 Transfer method.
// If the receiver for the tx is axelar's GMP address,
// Then it expects a payload of the gmptypes.MsgRelayPrice msg.
// If it does not have this format, it will error out.
// If it does, it will build a MsgTransfer with the payload.
func (k Keeper) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	gmpParams := k.gmpKeeper.GetParams(ctx)

	if msg.Receiver == gmpParams.GmpAddress {
		relayMsg := &gmptypes.MsgRelayPrice{}
		// safe byte conversion for invalid UTF-8
		bz := make([]byte, len(msg.Memo))
		copy(bz, msg.Memo)
		err := relayMsg.Unmarshal(bz)
		if err == nil {
			// If the payload is not a relayMsg type, then a user is trying to perform GMP
			// without the proper payload. This transaction be considered to be by a bad actor.
			k.Logger(ctx).With(err).Error("unexpected object while trying to relay data to GMP")
			return nil, err
		}

		gmpTransferMsg, err := k.gmpKeeper.BuildGmpRequest(goCtx, relayMsg)
		if err != nil {
			return nil, err
		}
		return k.Keeper.Transfer(goCtx, gmpTransferMsg)
	}
	return k.Keeper.Transfer(goCtx, msg)
}

// NewKeeper creates a new IBC transfer Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, paramSpace paramtypes.Subspace,
	ics4Wrapper porttypes.ICS4Wrapper, channelKeeper types.ChannelKeeper, portKeeper types.PortKeeper,
	authKeeper types.AccountKeeper, bankKeeper types.BankKeeper, scopedKeeper exported.ScopedKeeper,
	gmpKeeper GmpKeeper,
) Keeper {
	// ensure ibc transfer module account is set
	if addr := authKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the IBC transfer module account has not been set")
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		paramSpace:    paramSpace,
		ics4Wrapper:   ics4Wrapper,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		authKeeper:    authKeeper,
		bankKeeper:    bankKeeper,
		scopedKeeper:  scopedKeeper,
		gmpKeeper:     gmpKeeper,
	}
}
