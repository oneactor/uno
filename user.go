package uno

type User interface {
	GetId() int
	GetNick() string
	GetCards() map[int]Card
	GetCardsNum() int
	AddCard(Card)
	RemoveCard(int)
	RemoveCards([]int)
}
