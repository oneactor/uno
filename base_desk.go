package uno

import (
	"container/list"
)

type BaseDesk struct {
	cards map[int]Card // 卡组
	list  *list.List   // 顺序
}

func NewBaseDesk() Desk {
	return &BaseDesk{
		cards: map[int]Card{},
		list:  list.New(),
	}
}

func (this *BaseDesk) PrepareCards(cards []Card) {
	for _, v := range cards {
		this.cards[v.GetId()] = v
	}
}

func (this *BaseDesk) Shuffle() map[int]Card {
	for k, _ := range this.cards {
		this.list.PushBack(k)
	}
	return this.cards
}

func (this *BaseDesk) GetCards() map[int]Card {
	return this.cards
}

func (this *BaseDesk) GetCard(id int) (Card, bool) {
	card, ok := this.cards[id]
	if ok {
		return card, true
	} else {
		return nil, false
	}
}

func (this *BaseDesk) GetNext() Card {
	index := this.list.Front()
	this.list.Remove(index)
	card := this.cards[index.Value.(int)]
	return card
}

func (this *BaseDesk) GetLeftCount() int {
	return this.list.Len()
}
