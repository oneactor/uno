package blackjack21

import (
	"container/list"
	"github.com/jesusslim/uno"
)

type BlackJack21Desk struct {
	cards map[int]uno.Card
	list  *list.List
}

func NewBlackJack21Desk() uno.Desk {
	return &BlackJack21Desk{
		cards: map[int]uno.Card{},
		list:  list.New(),
	}
}

func NewDesk() uno.Desk {
	return NewBlackJack21Desk()
}

func (this *BlackJack21Desk) PrepareCards(cards []uno.Card) {
	for _, v := range cards {
		this.cards[v.GetId()] = v
	}
}

func (this *BlackJack21Desk) Shuffle() map[int]uno.Card {
	for k, _ := range this.cards {
		this.list.PushBack(k)
	}
	return this.cards
}

func (this *BlackJack21Desk) GetCards() map[int]uno.Card {
	return this.cards
}

func (this *BlackJack21Desk) GetCard(id int) (uno.Card, bool) {
	card, ok := this.cards[id]
	if ok {
		return card, true
	} else {
		return nil, false
	}
}

func (this *BlackJack21Desk) GetNext() uno.Card {
	index := this.list.Front()
	this.list.Remove(index)
	card := this.cards[index.Value.(int)]
	return card
}

func (this *BlackJack21Desk) GetLeftCount() int {
	return this.list.Len()
}
