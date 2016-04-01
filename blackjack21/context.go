package blackjack21

import (
	"fmt"
	"github.com/jesusslim/uno"
)

type Points struct {
	common_points int
	ace_num       int
}

const (
	ACTION_HIT       = 1
	ACTION_STAND     = 2
	ACTION_DOUBLE    = 3
	ACTION_SURRENDER = 4
)

func GetActions() map[int]string {
	return map[int]string{
		ACTION_HIT:       "HIT",
		ACTION_STAND:     "STAND",
		ACTION_DOUBLE:    "DOUBLE",
		ACTION_SURRENDER: "SURRENDER",
	}
}

type BlackJack21Context struct {
	uno.BaseContext
	points   map[int]*Points
	turn_end bool
}

func NewBlackJack21Context(desk uno.Desk) *BlackJack21Context {
	return &BlackJack21Context{
		*uno.NewBaseContext(desk),
		map[int]*Points{},
		false,
	}
}

func NewContext(desk uno.Desk) uno.Context {
	return NewBlackJack21Context(desk)
}

func (this *BlackJack21Context) AddUser(user uno.User) {
	this.BaseContext.AddUser(user)
	this.points[user.GetId()] = &Points{
		common_points: 0,
		ace_num:       0,
	}
}

func (this *BlackJack21Context) TurnNext() {
	this.BaseContext.TurnNext()
	this.turn_end = false
}

func (this *BlackJack21Context) OnPlay(ids []int, params map[string]interface{}) {
	action, ok := params["action"]
	if ok {
		user := this.GetNowUser()
		switch action {
		case ACTION_HIT, ACTION_DOUBLE:
			this.Draw(1)
			break
		case ACTION_STAND:
			this.turn_end = true
			fmt.Println("选择停牌")
			break
		case ACTION_SURRENDER:
			this.turn_end = true
			user.SetAttr("is_out", true)
			fmt.Println("放弃比赛")
			break
		}
	}
}

func (this *BlackJack21Context) ReCountPoints() {
	user := this.GetNowUser()
	if !user.GetAttrBool("isOut") {
		point := &Points{0, 0}
		for _, card := range user.GetCards() {
			switch card.GetTypeId() {
			case JACK_COMMON:
				point.common_points += card.GetAttrInt("points")
				break
			case JACK_10:
				point.common_points += 10
				break
			case JACK_ACE:
				point.ace_num++
				break
			}
		}
		this.points[user.GetId()] = point
	}
}

func (this *BlackJack21Context) CheckTurnEnd() bool {
	user := this.GetNowUser()
	if user.GetAttrBool("isOut") {
		delete(this.points, user.GetId())
	} else {
		this.ReCountPoints()
		point := this.points[user.GetId()]
		if point.common_points+point.ace_num > 21 {
			fmt.Println("爆牌")
			user.SetAttr("isOut", true)
			this.turn_end = true
			delete(this.points, user.GetId())
		}
	}
	if this.turn_end {
		fmt.Println("回合结束")
		return true
	}
	return false
}

func (this *BlackJack21Context) CheckWinner() bool {
	max_id := 0
	max := 0
	for id, point := range this.points {
		min := 0
		common := point.common_points
		min = common + point.ace_num
		for i := 0; i < point.ace_num; i++ {
			if min > 21 {
				break
			}
			if min == 21 {
				if min > max {
					max = min
					max_id = id
				}
				break
			}
			min = min + 9
		}
		if min <= 21 {
			if min > max {
				max = min
				max_id = id
			}
		}
	}
	users := this.GetUsers()
	if max_id > 0 {
		fmt.Println("Winner is ", users[max_id].GetNick())
		fmt.Println("His score:", fmt.Sprintf("%d", max))
	} else {
		fmt.Println("No winner.")
	}
	return true
}
