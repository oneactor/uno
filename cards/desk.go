package cards

import (
	"container/list"
	"github.com/jesusslim/uno"
	"math/rand"
	"time"
)

type UnoDesk struct {
	cards []uno.Card
	list  *list.List
}

func NewDesk() uno.Desk {
	return &UnoDesk{
		cards: []uno.Card{},
		list:  list.New(),
	}
}

func (this *UnoDesk) PrepareCards(cards []uno.Card) {
	for _, v := range cards {
		this.cards = append(this.cards, v)
	}
}

func (this *UnoDesk) Shuffle() []uno.Card {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(this.cards)
	temp_arr := []int{}
	for k, _ := range this.cards {
		temp_arr = append(temp_arr, k)
	}
	for i := length - 1; i > 0; i-- {
		r := seed.Intn(length)
		temp_arr[r], temp_arr[i] = temp_arr[i], temp_arr[r]
	}
	for _, v := range temp_arr {
		this.list.PushBack(v)
	}
	return this.cards
}

func (this *UnoDesk) GetCards() []uno.Card {
	return this.cards
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
