package uno

/**
 * 上下文接口
 */
type Context interface {
	AddUser(User)                         // 一局添加用户
	GetUsers() map[int]User               // 获取全部用户
	GetNextUser() User                    // 获取下家
	GetNowUser() User                     // 当前出牌者
	IsClockwise() bool                    // 是否顺时针
	SetClockwise(bool) bool               // 设置轮转方向
	GetCardsUsed() map[int]Card           // 获取已被使用的卡
	GetCardsUsedQueue() []int             // 获取已被使用的卡顺序
	GetCardsNow() map[int]Card            // 获取当前回合出的牌
	GetCardsLast() map[int]Card           // 获取上回合出的牌
	CardsUse(map[int]Card)                // 使用卡牌
	CardsUseWithId(...int)                // 使用卡牌
	TurnNext()                            // 流转至下一回合
	OnStart()                             // 起始时动作
	CheckPlay([]int) (bool, string)       // 检查出牌是否允许
	OnPlay([]int, map[string]interface{}) // 出牌
	Draw(int) (bool, string)              // 摸牌
	OnEnd() (bool, string)                // 回合结束动作
}
