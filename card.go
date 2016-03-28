package uno

type Card interface {
	GetId() int
	GetTypeId() int
	GetAttr(string) interface{}
	GetAttrInt(string) int
	GetAttrStr(string) string
	OnDraw()
}
