package blackjack21

import (
	"github.com/jesusslim/uno"
)

type BlackJack21Desk struct {
	uno.BaseDesk
}

func NewBlackJack21Desk() uno.Desk {
	return uno.NewBaseDesk()
}

func NewDesk() uno.Desk {
	return NewBlackJack21Desk()
}
