package uno

type Context interface {
	AddUser(User)
	GetUsers() map[int]User
	GetNextUser() User
	GetNowUser() User
	IsClockwise() bool
	SetClockwise(bool) bool
	GetCardsUsed() map[int]Card
	GetCardsUsedQueue() []int
	GetCardsNow() map[int]Card
	GetCardsLast() map[int]Card
	CardsUse(map[int]Card)
	CardsUseWithId(...int)
	TurnNext()
	OnStart()
	CheckPlay([]int) (bool, string)
	OnPlay([]int, map[string]interface{})
	Draw(int) (bool, string)
	OnEnd() (bool, string)
}
