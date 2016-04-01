package cards

import (
	"github.com/jesusslim/uno"
)

/**
 * uno卡组
 */
type UnoDesk struct {
	uno.BaseDesk
}

func NewDesk() uno.Desk {
	return uno.NewBaseDesk()
}
