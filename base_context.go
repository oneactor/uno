package uno

import (
	"fmt"
	"math/rand"
	"time"
)

type BaseContext struct {
	users            map[int]User // 玩家/用户
	clockwise        bool         // 顺时针
	user_queue       []int        // 用户顺序
	user_index_now   int          // 当前用户序号
	cards_used       map[int]Card // 已被使用的卡牌
	cards_used_queue []int        // 使用顺序
	cards_last       map[int]Card // 上回合使用的牌
	cards_now        map[int]Card // 这回合使用的牌
	desk             Desk         // 牌库
	seed             *rand.Rand   // 随机数种子
}

func NewBaseContext(desk Desk) *BaseContext {
	return &BaseContext{
		users:            map[int]User{},
		clockwise:        true,
		user_queue:       []int{},
		user_index_now:   0,
		cards_used:       map[int]Card{},
		cards_used_queue: []int{},
		cards_last:       map[int]Card{},
		cards_now:        map[int]Card{},
		desk:             desk,
		seed:             rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func NewBaseContextInterface(desk Desk) Context {
	return NewBaseContext(desk)
}

func (this *BaseContext) AddUser(user User) {
	this.users[user.GetId()] = user
	this.user_queue = append(this.user_queue, user.GetId())
}

func (this *BaseContext) GetUsers() map[int]User {
	return this.users
}

func (this *BaseContext) GetNextUser() User {
	var next_index int
	if this.clockwise {
		next_index = (this.user_index_now + 1) % len(this.user_queue)
	} else {
		next_index = (this.user_index_now - 1) % len(this.user_queue)
	}
	return this.users[this.user_queue[next_index]]
}

func (this *BaseContext) GetNowUser() User {
	return this.users[this.user_queue[this.user_index_now]]
}

func (this *BaseContext) GetCardsUsed() map[int]Card {
	return this.cards_used
}

func (this *BaseContext) GetCardsUsedQueue() []int {
	return this.cards_used_queue
}

func (this *BaseContext) GetCardsNow() map[int]Card {
	return this.cards_now
}

func (this BaseContext) SetCardsNow(cards map[int]Card) {
	this.cards_now = cards
}

func (this *BaseContext) GetCardsLast() map[int]Card {
	return this.cards_last
}

func (this *BaseContext) SetCardsLast(cards map[int]Card) {
	this.cards_last = cards
}

func (this *BaseContext) CardsUse(cards map[int]Card) {
	for k, v := range cards {
		this.cards_now[k] = v
		this.cards_used[k] = v
		this.cards_used_queue = append(this.cards_used_queue, k)
	}
}

func (this *BaseContext) CardsUseWithId(ids ...int) {
	for _, id := range ids {
		card, ok := this.desk.GetCard(id)
		if ok {
			this.cards_now[id] = card
			this.cards_used[id] = card
			this.cards_used_queue = append(this.cards_used_queue, id)
		}
	}
}

func (this *BaseContext) TurnNext() {
	fmt.Println("Turn Next...")
	if this.clockwise {
		this.user_index_now = (this.user_index_now + 1) % len(this.user_queue)
	} else {
		temp := this.user_index_now - 1
		if temp < 0 {
			temp = temp + len(this.user_queue)
		}
		this.user_index_now = temp % len(this.user_queue)
	}
}

func (this *BaseContext) GetDesk() Desk {
	return this.desk
}

func (this *BaseContext) OnStart() {
	fmt.Println("On Start...")
}

func (this *BaseContext) OnEnd() (bool, string) {
	if this.GetDesk().GetLeftCount() == 0 {
		return false, "Cards not enough"
	}
	return true, ""
}

func (this *BaseContext) OnPlay(ids []int, params map[string]interface{}) {
	fmt.Println("On Play...")
}

func (this *BaseContext) Draw(num int) (bool, string) {
	if this.GetDesk().GetLeftCount() < num {
		return false, "Cards not enough."
	} else {
		u := this.GetNowUser()
		fmt.Println(u.GetNick(), " Draw ", fmt.Sprintf("%d", num), " cards.")
		for i := 0; i < num; i++ {
			card := this.GetDesk().GetNext()
			fmt.Println(card)
			u.AddCard(card)
		}
	}
	return true, ""
}

func (this *BaseContext) CheckPlay(ids []int) (bool, string) {
	return true, ""
}

func (this *BaseContext) IsClockwise() bool {
	return this.clockwise
}

func (this *BaseContext) SetClockwise(clockwise bool) bool {
	this.clockwise = clockwise
	return this.clockwise
}

func (this *BaseContext) GetRandom(base int) int {
	return this.seed.Intn(base)
}
