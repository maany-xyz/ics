package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Ensure MsgCustomProposal implements sdk.Msg
var _ sdk.Msg = &MsgConsumerGovProposal{}

// MsgCustomProposal defines a governance proposal message
func (m *MsgConsumerGovProposal) Route() string {
	return RouterKey
}

func (m *MsgConsumerGovProposal) Type() string {
	return "ConsumerGovProposal"
}

// ValidateBasic performs basic validation checks
func (m *MsgConsumerGovProposal) ValidateBasic() error {
	if len(m.Title) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "title cannot be empty")
	}
	if len(m.Description) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "description cannot be empty")
	}
	// Ensure authority is a valid address
	_, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid authority address")
	}
	return nil
}

// GetSigners returns the expected signers for MsgCustomProposal
func (m *MsgConsumerGovProposal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err) // Should never happen since ValidateBasic checks this
	}
	return []sdk.AccAddress{addr}
}