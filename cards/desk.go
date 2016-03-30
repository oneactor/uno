package cards

import (
	"container/list"
	"github.com/jesusslim/uno"
)

/**
 * uno卡组
 */
type UnoDesk struct {
	cards map[int]uno.Card // 卡组
	list  *list.List       // 顺序
}

func NewDesk() uno.Desk {
	return &UnoDesk{
		cards: map[int]uno.Card{},
		list:  list.New(),
	}
}

func (this *UnoDesk) PrepareCards(cards []uno.Card) {
	for _, v := range cards {
		this.cards[v.GetId()] = v
	}
}

func (this *UnoDesk) Shuffle() map[int]uno.Card {
	for k, _ := range this.cards {
		this.list.PushBack(k)
	}
	return this.cards
}

func (this *UnoDesk) GetCards() map[int]uno.Card {
	return this.cards
}

func (this *UnoDesk) GetCard(id int) (uno.Card, bool) {
	card, ok := this.cards[id]
	if ok {
		return card, true
	} else {
		return nil, false
	}
}

func (this *UnoDesk) GetNext() uno.Card {
	index := this.list.Front()
	this.list.Remove(index)
	card := this.cards[index.Value.(int)]
	return card
}

func (this *UnoDesk) GetLeftCount() int {
	return this.list.Len()
}
