package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimOwner = "claim_owner"

var _ sdk.Msg = &MsgClaimOwner{}

func NewMsgClaimOwner(creator string) *MsgClaimOwner {
	return &MsgClaimOwner{
		Creator: creator,
	}
}

func (msg *MsgClaimOwner) Route() string {
	return RouterKey
}

func (msg *MsgClaimOwner) Type() string {
	return TypeMsgClaimOwner
}

func (msg *MsgClaimOwner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
