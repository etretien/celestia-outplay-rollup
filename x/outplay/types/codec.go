package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateProfile{}, "outplay/CreateProfile", nil)
	cdc.RegisterConcrete(&MsgCreateChallenge{}, "outplay/CreateChallenge", nil)
	cdc.RegisterConcrete(&MsgAcceptChallenge{}, "outplay/AcceptChallenge", nil)
	cdc.RegisterConcrete(&MsgDeclineChallenge{}, "outplay/DeclineChallenge", nil)
	cdc.RegisterConcrete(&MsgSubmitScore{}, "outplay/SubmitScore", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProfile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChallenge{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAcceptChallenge{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeclineChallenge{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitScore{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
