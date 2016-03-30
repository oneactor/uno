package uno

/**
 * 卡牌接口
 */
type Card interface {
	GetId() int                 //id
	GetTypeId() int             //类型id
	GetAttr(string) interface{} //获取某个属性
	GetAttrInt(string) int
	GetAttrStr(string) string
	OnDraw() //摸牌触发
}
