package uno

type Desk interface {
	PrepareCards([]Card)
	Shuffle() []Card
	GetCards() []Card
	GetNext() Card
	GetLeftCount() int
}
