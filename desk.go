package uno

type Desk interface {
	PrepareCards([]Card)
	Shuffle() map[int]Card
	GetCards() map[int]Card
	GetCard(int) (Card, bool)
	GetNext() Card
	GetLeftCount() int
}
