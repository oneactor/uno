package uno

type User interface {
	GetId() int
	GetNick() string
	GetCards() []Card
	AddCard(Card)
	RemoveCard(Card)
	AddCards([]Card)
	RemoveCards([]Card)
}
