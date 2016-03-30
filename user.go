package uno

type User interface {
	GetId() int
	GetNick() string
	GetCards() map[int]Card
	GetCardsNum() int
	AddCard(Card)
	RemoveCard(int)
	RemoveCards([]int)
	SetAttr(string, interface{}) bool
	GetAttr(string) (interface{}, bool)
	GetAttrStr(string) string
	GetAttrBool(string) bool
}
