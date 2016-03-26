package uno

type Card interface {
	GetId() int
	GetTypeId() int
	OnDraw()
	CheckPlay() (bool, string)
	OnPlay()
}
