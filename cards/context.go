package cards

import (
	"fmt"
	"github.com/jesusslim/uno"
)

type UnoContext struct {
	users            map[int]uno.User
	clockwise        bool
	user_queue       []int
	user_index_now   int
	cards_used       map[int]uno.Card
	cards_used_queue []int
	cards_last       map[int]uno.Card
	cards_now        map[int]uno.Card
	desk             uno.Desk
	color_tmp        int
	draw_num_ext     int
	this_turn_played bool
}

func NewUnoContext(desk uno.Desk) *UnoContext {
	return &UnoContext{
		users:            map[int]uno.User{},
		clockwise:        true,
		user_queue:       []int{},
		user_index_now:   0,
		cards_used:       map[int]uno.Card{},
		cards_used_queue: []int{},
		cards_last:       map[int]uno.Card{},
		cards_now:        map[int]uno.Card{},
		desk:             desk,
		color_tmp:        COLOR_BLACK,
		draw_num_ext:     0,
		this_turn_played: false,
	}
}

func NewContext(desk uno.Desk) uno.Context {
	return NewUnoContext(desk)
}

func (this *UnoContext) AddUser(user uno.User) {
	this.users[user.GetId()] = user
	this.user_queue = append(this.user_queue, user.GetId())
}

func (this *UnoContext) GetUsers() map[int]uno.User {
	return this.users
}

func (this *UnoContext) GetNextUser() uno.User {
	var next_index int
	if this.clockwise {
		next_index = (this.user_index_now + 1) % len(this.user_queue)
	} else {
		next_index = (this.user_index_now - 1) % len(this.user_queue)
	}
	return this.users[this.user_queue[next_index]]
}

func (this *UnoContext) GetNowUser() uno.User {
	return this.users[this.user_queue[this.user_index_now]]
}

func (this *UnoContext) IsClockwise() bool {
	return this.clockwise
}

func (this *UnoContext) SetClockwise(clockwise bool) bool {
	this.clockwise = clockwise
	return this.clockwise
}

func (this *UnoContext) GetCardsUsed() map[int]uno.Card {
	return this.cards_used
}

func (this *UnoContext) GetCardsUsedQueue() []int {
	return this.cards_used_queue
}

func (this *UnoContext) GetCardsNow() map[int]uno.Card {
	return this.cards_now
}

func (this *UnoContext) GetCardsLast() map[int]uno.Card {
	return this.cards_last
}

func (this *UnoContext) CardsUse(cards map[int]uno.Card) {
	for k, v := range cards {
		this.cards_now[k] = v
		this.cards_used[k] = v
		this.cards_used_queue = append(this.cards_used_queue, k)
	}
}

func (this *UnoContext) CardsUseWithId(ids ...int) {
	for _, id := range ids {
		card, ok := this.desk.GetCard(id)
		if ok {
			this.cards_now[id] = card
			this.cards_used[id] = card
			this.cards_used_queue = append(this.cards_used_queue, id)
		}
	}
}

func (this *UnoContext) TurnNext() {
	this.cards_last = this.cards_now
	this.cards_now = map[int]uno.Card{}
	if this.clockwise {
		this.user_index_now = (this.user_index_now + 1) % len(this.user_queue)
	} else {
		temp := this.user_index_now - 1
		if temp < 0 {
			temp = temp + len(this.user_queue)
		}
		this.user_index_now = temp % len(this.user_queue)
	}
	fmt.Println("user_index_now:", this.user_index_now)
}

func (this *UnoContext) GetDesk() uno.Desk {
	return this.desk
}

func (this *UnoContext) OnStart() {

}

func (this *UnoContext) OnEnd() (bool, string) {
	if !this.this_turn_played {
		//draw
		num := this.draw_num_ext
		if num == 0 {
			num = 1
		}
		ok, _ := this.Draw(num)
		if !ok {
			//game end
			return false, "Cards not enough."
		}
	}
	this.TurnNext()
	return true, ""
}

func (this *UnoContext) OnPlay(ids []int, params map[string]interface{}) {
	u := this.GetNowUser()
	for k, id := range ids {
		if k == 0 {
			card, _ := this.GetDesk().GetCard(id)
			if card.GetAttrInt("color") == COLOR_BLACK {
				color_in_param, ok := params["color"]
				if ok {
					this.color_tmp = color_in_param.(int)
				}
			} else {
				this.color_tmp = card.GetAttrInt("color")
			}
		}
		u.RemoveCard(id)
		this.CardsUseWithId(id)
	}
}

func (this *UnoContext) Draw(num int) (bool, string) {
	if this.GetDesk().GetLeftCount() < num {
		return false, "Cards not enough."
	} else {
		u := this.GetNowUser()
		for i := 0; i < num; i++ {
			card := this.GetDesk().GetNext()
			u.AddCard(card)
		}
	}
	this.draw_num_ext = 0
	return true, ""
}

/**
 * 出牌检查
 * @param  ids []int 出牌id数组
 * @return bool,string
 */
func (this *UnoContext) CheckPlay(ids []int) (bool, string) {
	if len(ids) == 0 {
		return false, "No cards."
	}
	//check used
	used := this.GetCardsUsed()
	for _, id := range ids {
		_, ok := used[id]
		if ok {
			return false, "Card " + fmt.Sprintf("%d", id) + " is used."
		}
	}
	cards := this.GetDesk().GetCards()
	if len(ids) > 1 {
		//大于一张 判断是否相同
		var points int
		var color int
		for k, id := range ids {
			card := cards[id]
			if card.GetAttrInt("extType") != TYPE_COMMON {
				return false, "Card type error."
			}
			if k == 0 {
				points = card.GetAttrInt("points")
				color = card.GetAttrInt("color")
			} else {
				if (points != card.GetAttrInt("points")) || (color != card.GetAttrInt("color")) {
					return false, "Cards not same."
				}
			}
		}
	}
	card_example := cards[ids[0]]
	the_points := card_example.GetAttrInt("points")
	the_color := card_example.GetAttrInt("color")
	the_type_id := card_example.GetAttrInt("typeId")
	the_ext_type := card_example.GetAttrInt("extType")
	cards_last := this.GetCardsLast()
	for _, v := range cards_last {
		switch v.GetAttrInt("extType") {
		case TYPE_COMMON:
			switch the_ext_type {
			case TYPE_COMMON:
				if (v.GetAttrInt("color") != the_color) && (v.GetAttrInt("points") != the_points) {
					return false, "花色且点数不对"
				}
				break
			case TYPE_USEAGE:
				if (v.GetAttrInt("color") != the_color) && (v.GetAttrInt("typeId") != the_type_id) {
					return false, "功能牌花色或类型不对"
				}
				break
			case TYPE_ALL_CAN:
				if the_type_id == CARD_WILD_DRAW_4 {
					u := this.GetNowUser()
					cards_in_hands := u.GetCards()
					for k, i_h := range cards_in_hands {
						if k == card_example.GetId() {
							continue
						}
						//手中无同花色可出时才允许
						if (i_h.GetAttrInt("color") == v.GetAttrInt("color")) && (i_h.GetAttrInt("typeId") != CARD_WILD_DRAW_4) {
							return false, "手中有同花色可出牌,不允许+4."
						}
					}
				}
				break
			}
			break
		case TYPE_USEAGE:
			switch the_ext_type {
			case TYPE_COMMON:
				return false, "只允许跳过"
				break
			case TYPE_USEAGE:
				if the_type_id != CARD_JUMP {
					return false, "只允许跳过"
				}
				break
			case TYPE_ALL_CAN:
				return false, "只允许跳过"
				break
			}
			break
		case TYPE_ALL_CAN:
			if v.GetAttrInt("typeId") == CARD_WILD {
				switch the_ext_type {
				case TYPE_COMMON:
					if this.color_tmp != the_color {
						return false, "花色不对"
					}
					break
				case TYPE_USEAGE:
					if this.color_tmp != the_color {
						return false, "功能牌花色或类型不对"
					}
					break
				case TYPE_ALL_CAN:
					if the_type_id == CARD_WILD_DRAW_4 {
						u := this.GetNowUser()
						cards_in_hands := u.GetCards()
						for k, i_h := range cards_in_hands {
							if k == card_example.GetId() {
								continue
							}
							//手中无同花色可出时才允许
							if (i_h.GetAttrInt("color") == this.color_tmp) && (i_h.GetAttrInt("typeId") != CARD_WILD_DRAW_4) {
								return false, "手中有同花色可出牌,不允许+4."
							}
						}
					}
					break
				}
			} else if v.GetAttrInt("typeId") == CARD_WILD_DRAW_4 {
				if the_type_id == CARD_REV {
					if the_color != this.color_tmp {
						return false, "颜色错误"
					}
				} else if the_type_id == CARD_WILD_DRAW_4 {
					//ok
				} else {
					return false, "出牌错误"
				}
			}
			break
		}
		//只做一次
		break
	}
	//判断反转/摸牌
	if the_type_id == CARD_REV {
		this.SetClockwise(!this.IsClockwise())
	} else if the_type_id == CARD_DRAW_2 {
		this.draw_num_ext += 2
	} else if the_type_id == CARD_WILD_DRAW_4 {
		this.draw_num_ext += 4
	}
	this.this_turn_played = true
	return true, ""
}

func (this *UnoContext) CheckWinner() bool {
	if this.GetNowUser().GetCardsNum() == 0 {
		return true
	} else {
		return false
	}
}

func (this *UnoContext) CheckUno() bool {
	if this.GetNowUser().GetCardsNum() == 1 {
		return true
	} else {
		return false
	}
}
