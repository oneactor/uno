package blackjack21

import (
	"github.com/jesusslim/uno"
)

type User struct {
	uno.BaseUser
	is_out bool //是否出局标志
}

func NewUser(id int, nick string) uno.User {
	return &User{
		*uno.NewBaseUser(id, nick),
		false,
	}
}

func (this *User) GetAttrBool(name string) bool {
	switch name {
	case "is_out", "isOut":
		return this.is_out
		break
	}
	return false
}
