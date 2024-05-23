package runes


import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgRune struct {
	Sender sdk.AccAddress `json: "sender"`
	RuneData string       `json: "rune_data"`
}

func NewMsgRune(sender sdk.AccAddress, runeData string) MsgRune {
	return MsgRune{
		Sender: sender,
		RuneData: runeData,
	}
}

func (msg MsgRune) Route() string {return "runes"}

func (msg MsgRune) Type() string {return "create_rune"}

