package uno

type BaseUser struct {
	id    int
	nick  string
	cards map[int]Card
}

func NewBaseUser(id int, nick string) *BaseUser {
	return &BaseUser{
		id:    id,
		nick:  nick,
		cards: map[int]Card{}, // 手牌
	}
}

func NewBaseUserInterface(id int, nick string) User {
	return NewBaseUser(id, nick)
}

func (this *BaseUser) GetId() int {
	return this.id
}

func (this *BaseUser) GetNick() string {
	return this.nick
}

func (this *BaseUser) GetCards() map[int]Card {
	return this.cards
}

func (this *BaseUser) GetCardsNum() int {
	return len(this.cards)
}

func (this *BaseUser) AddCard(card Card) {
	this.cards[card.GetId()] = card
}

func (this *BaseUser) RemoveCard(id int) {
	delete(this.cards, id)
}

func (this *BaseUser) RemoveCards(ids []int) {
	for _, id := range ids {
		delete(this.cards, id)
	}
}

func (this *BaseUser) SetAttr(name string, value interface{}) bool {
	return false
}

func (this *BaseUser) GetAttr(name string) (interface{}, bool) {
	return nil, false
}

func (this *BaseUser) GetAttrStr(name string) string {
	return ""
}

func (this *BaseUser) GetAttrBool(name string) bool {
	return false
}
