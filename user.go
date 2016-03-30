package uno

/**
 * 用户接口
 */
type User interface {
	GetId() int                         // 获取id
	GetNick() string                    // 获取昵称
	GetCards() map[int]Card             // 手牌
	GetCardsNum() int                   // 手牌数
	AddCard(Card)                       // 加牌
	RemoveCard(int)                     // 去牌
	RemoveCards([]int)                  // 去牌批量
	SetAttr(string, interface{}) bool   // 设置属性
	GetAttr(string) (interface{}, bool) // 获取属性
	GetAttrStr(string) string
	GetAttrBool(string) bool
}
