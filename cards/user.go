package cards

import (
	"github.com/jesusslim/uno"
)

type UnoUser struct {
	id    int
	nick  string
	cards map[int]uno.Card
}

func NewUnoUser(id int, nick string) uno.User {
	return &UnoUser{
		id:    id,
		nick:  nick,
		cards: map[int]uno.Card{},
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
