package cards

import (
	"github.com/jesusslim/uno"
)

/**
 * uno用户
 */
type UnoUser struct {
	uno.BaseUser
	is_uno bool
}

func NewUnoUserStruct(id int, nick string) *UnoUser {
	return &UnoUser{
		*uno.NewBaseUser(id, nick),
		false, // 是否uno (只剩一张手牌)
	}
}

func NewUnoUser(id int, nick string) uno.User {
	return NewUnoUserStruct(id, nick)
}

func (this *UnoUser) SetAttr(name string, value interface{}) bool {
	switch name {
	case "isUno", "is_nuo":
		this.is_uno = value.(bool)
		return true
		break
	}
	return false
}

func (this *UnoUser) GetAttr(name string) (interface{}, bool) {
	switch name {
	case "isUno", "is_nuo":
		return this.is_uno, true
		break
	}
	return nil, false
}

func (this *UnoUser) GetAttrStr(name string) string {
	return ""
}

func (this *UnoUser) GetAttrBool(name string) bool {
	switch name {
	case "isUno", "is_nuo":
		return this.is_uno
		break
	}
	return false
}
