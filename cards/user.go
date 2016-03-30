package cards

import (
	"github.com/jesusslim/uno"
)

/**
 * uno用户
 */
type UnoUser struct {
	id     int
	nick   string
	cards  map[int]uno.Card
	is_uno bool
}

func NewUnoUser(id int, nick string) uno.User {
	return &UnoUser{
		id:     id,
		nick:   nick,
		cards:  map[int]uno.Card{}, // 手牌
		is_uno: false,              // 是否uno (只剩一张手牌)
	}
}

func (this *UnoUser) GetId() int {
	return this.id
}

func (this *UnoUser) GetNick() string {
	return this.nick
}

func (this *UnoUser) GetCards() map[int]uno.Card {
	return this.cards
}

func (this *UnoUser) GetCardsNum() int {
	return len(this.cards)
}

func (this *UnoUser) AddCard(card uno.Card) {
	this.cards[card.GetId()] = card
}

func (this *UnoUser) RemoveCard(id int) {
	delete(this.cards, id)
}

func (this *UnoUser) RemoveCards(ids []int) {
	for _, id := range ids {
		delete(this.cards, id)
	}
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
