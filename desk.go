package uno

/**
 * 卡组接口
 */
type Desk interface {
	PrepareCards([]Card)      // 加入卡牌
	Shuffle() map[int]Card    // 洗牌
	GetCards() map[int]Card   // 获取卡牌
	GetCard(int) (Card, bool) // 获取单张卡牌
	GetNext() Card            // 获取下一张卡牌
	GetLeftCount() int        // 获取卡牌剩余总数
}
