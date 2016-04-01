package blackjack21

import (
	"fmt"
	"github.com/jesusslim/uno"
	"math/rand"
	"time"
)

type Points struct {
	common_points int
	ace_num       int
}

type BlackJack21Context struct {
	users            map[int]uno.User // 玩家/用户
	clockwise        bool             // 顺时针
	user_queue       []int            // 用户顺序
	user_index_now   int              // 当前用户序号
	cards_used       map[int]uno.Card // 已被使用的卡牌
	cards_used_queue []int            // 使用顺序
	cards_last       map[int]uno.Card // 上回合使用的牌
	cards_now        map[int]uno.Card // 这回合使用的牌
	desk             uno.Desk         // 牌库
	seed             *rand.Rand       // 随机数种子
	points           map[int]*Points
}

func NewBlackJack21Context(desk uno.Desk) *BlackJack21Context {
	return &BlackJack21Context{
		users:            map[int]uno.User{},
		clockwise:        true,
		user_queue:       []int{},
		user_index_now:   0,
		cards_used:       map[int]uno.Card{},
		cards_used_queue: []int{},
		cards_last:       map[int]uno.Card{},
		cards_now:        map[int]uno.Card{},
		desk:             desk,
		seed:             rand.New(rand.NewSource(time.Now().UnixNano())),
		points:           map[int]*Points{},
	}
}

func NewContext(desk uno.Desk) uno.Context {
	return NewBlackJack21Context(desk)
}

func (this *BlackJack21Context) AddUser(user uno.User) {
	this.users[user.GetId()] = user
	this.points[user.GetId()] = &Points{
		common_points: 0,
		ace_num:       0,
	}
	this.user_queue = append(this.user_queue, user.GetId())
}

func (this *BlackJack21Context) GetUsers() map[int]uno.User {
	return this.users
}

func (this *BlackJack21Context) GetNextUser() uno.User {
	var next_index int
	next_index = (this.user_index_now + 1) % len(this.user_queue)
	return this.users[this.user_queue[next_index]]
}

func (this *BlackJack21Context) GetNowUser() uno.User {
	return this.users[this.user_queue[this.user_index_now]]
}

func (this *BlackJack21Context) GetCardsUsed() map[int]uno.Card {
	return this.cards_used
}

func (this *BlackJack21Context) GetCardsUsedQueue() []int {
	return this.cards_used_queue
}

func (this *BlackJack21Context) GetCardsNow() map[int]uno.Card {
	return this.cards_now
}

func (this *BlackJack21Context) GetCardsLast() map[int]uno.Card {
	return this.cards_last
}

func (this *BlackJack21Context) CardsUse(cards map[int]uno.Card) {
	for k, v := range cards {
		this.cards_now[k] = v
		this.cards_used[k] = v
		this.cards_used_queue = append(this.cards_used_queue, k)
	}
}

func (this *BlackJack21Context) CardsUseWithId(ids ...int) {
	for _, id := range ids {
		card, ok := this.desk.GetCard(id)
		if ok {
			this.cards_now[id] = card
			this.cards_used[id] = card
			this.cards_used_queue = append(this.cards_used_queue, id)
		}
	}
}

func (this *BlackJack21Context) TurnNext() {
	this.cards_last = this.cards_now
	this.cards_now = map[int]uno.Card{}
	this.user_index_now = (this.user_index_now + 1) % len(this.user_queue)
}

func (this *BlackJack21Context) GetDesk() uno.Desk {
	return this.desk
}

func (this *BlackJack21Context) OnStart() {
	this.Draw(1)
}

func (this *BlackJack21Context) OnEnd() (bool, string) {
	if this.GetDesk().GetLeftCount() == 0 {
		return false, "Cards not enough"
	}
	this.TurnNext()
	return true, ""
}

func (this *BlackJack21Context) OnPlay(ids []int, params map[string]interface{}) {
	user := this.GetNowUser()
	for _, card := range user.GetCards() {
		//TODO
	}
}

func (this *BlackJack21Context) Draw(num int) (bool, string) {
	if this.GetDesk().GetLeftCount() < num {
		return false, "Cards not enough."
	} else {
		u := this.GetNowUser()
		fmt.Println(u.GetNick(), " 摸 ", fmt.Sprintf("%d", num), " 张")
		for i := 0; i < num; i++ {
			card := this.GetDesk().GetNext()
			u.AddCard(card)
		}
	}
	return true, ""
}

func (this *BlackJack21Context) CheckPlay(ids []int) (bool, string) {
	return true, ""
}

func (this *BlackJack21Context) CheckWinner() bool {
	return false
}

func (this *BlackJack21Context) IsClockwise() bool {
	return this.clockwise
}

func (this *BlackJack21Context) SetClockwise(clockwise bool) bool {
	this.clockwise = clockwise
	return this.clockwise
}
