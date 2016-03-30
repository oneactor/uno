package cards

import (
	"fmt"
	"github.com/jesusslim/uno"
	"math/rand"
	"time"
)

//定义回合类型
const (
	TURN_TYPE_COMMON = 1 // 正常
	TURN_TYPE_DRAW   = 2 // 摸牌
	TURN_TYPE_JUMP   = 3 // 跳过
)

/**
 * uno上下文
 */
type UnoContext struct {
	users            map[int]uno.User // 玩家/用户
	clockwise        bool             // 顺时针
	user_queue       []int            // 用户顺序
	user_index_now   int              // 当前用户序号
	cards_used       map[int]uno.Card // 已被使用的卡牌
	cards_used_queue []int            // 使用顺序
	cards_last       map[int]uno.Card // 上回合使用的牌
	cards_now        map[int]uno.Card // 这回合使用的牌
	desk             uno.Desk         // 牌库
	this_turn_played bool             // 该回合是否成功出了牌
	seed             *rand.Rand       // 随机数种子

	//状态与回合相关
	color_tmp    int // 颜色
	points_tmp   int // 点数
	draw_num_ext int // 摸牌数
	turn_type    int // 回合类型
}

//根据出牌决定下回合类型
func (this *UnoContext) GetTurnTypeByCard(card_type_id int) int {
	switch card_type_id {
	case CARD_COMMON, CARD_WILD:
		return TURN_TYPE_COMMON
		break
	case CARD_JUMP:
		return TURN_TYPE_JUMP
		break
	case CARD_REV:
		return this.turn_type
		break
	case CARD_DRAW_2, CARD_WILD_DRAW_4:
		return TURN_TYPE_DRAW
		break
	}
	return this.turn_type
}

//新建uno上下文
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
		this_turn_played: false,
		seed:             rand.New(rand.NewSource(time.Now().UnixNano())),

		color_tmp:    COLOR_NO_MATCH,
		points_tmp:   POINTS_NO_MATCH,
		draw_num_ext: 0,
		turn_type:    TURN_TYPE_COMMON,
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
	if this.this_turn_played {
		this.cards_last = this.cards_now
		this.cards_now = map[int]uno.Card{}
	}
	if this.clockwise {
		this.user_index_now = (this.user_index_now + 1) % len(this.user_queue)
	} else {
		temp := this.user_index_now - 1
		if temp < 0 {
			temp = temp + len(this.user_queue)
		}
		this.user_index_now = temp % len(this.user_queue)
	}
	this.this_turn_played = false
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
			//turn status
			card, _ := this.GetDesk().GetCard(id)
			c_color := card.GetAttrInt("color")
			c_points := card.GetAttrInt("points")
			switch card.GetTypeId() {
			case CARD_COMMON:
				this.color_tmp = c_color
				this.points_tmp = c_points
				break
			case CARD_JUMP:
				break
			case CARD_REV:
				break
			case CARD_WILD, CARD_WILD_DRAW_4:
				this.points_tmp = POINTS_NO_MATCH
				color_in_param, ok := params["color"]
				if ok {
					this.color_tmp = color_in_param.(int)
				} else {
					colors := []int{COLOR_BLUE, COLOR_GREEN, COLOR_RED, COLOR_YELLOW}
					index := this.seed.Intn(len(colors))
					this.color_tmp = colors[index]
				}
				fmt.Println("指定颜色:", ConvertColor(this.color_tmp))
				break
			case CARD_DRAW_2:
				this.color_tmp = c_color
				this.points_tmp = POINTS_NO_MATCH
				break
			}
			this.turn_type = this.GetTurnTypeByCard(card.GetTypeId())
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
		fmt.Println(u.GetNick(), " 摸 ", fmt.Sprintf("%d", num), " 张")
		for i := 0; i < num; i++ {
			card := this.GetDesk().GetNext()
			u.AddCard(card)
		}
	}
	//重置回合类型
	this.turn_type = TURN_TYPE_COMMON
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
	//check used/in hands
	cards_in_hand := this.GetNowUser().GetCards()
	used := this.GetCardsUsed()
	for _, id := range ids {
		_, ok := used[id]
		if ok {
			return false, "Card " + fmt.Sprintf("%d", id) + " is used."
		}
		_, ok2 := cards_in_hand[id]
		if !ok2 {
			return false, "Card " + fmt.Sprintf("%d", id) + " is not in your hands."
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
	//the_ext_type := card_example.GetAttrInt("extType")
	cards_last := this.GetCardsLast()
	//ins
	switch this.turn_type {
	case TURN_TYPE_COMMON:
		switch the_type_id {
		case CARD_COMMON:
			if !this.Match(the_color, the_points, false) {
				return false, "花色且点数不对"
			}
			break
		case CARD_JUMP, CARD_REV, CARD_DRAW_2:
			if !this.Match(the_color, the_points, true) {
				return false, "功能牌花色不对"
			}
			break
		case CARD_WILD:
			break
		case CARD_WILD_DRAW_4:
			u := this.GetNowUser()
			cards_in_hands := u.GetCards()
			for k, i_h := range cards_in_hands {
				if k == card_example.GetId() {
					continue
				}
				//手中无同花色可出时才允许
				if (this.Match(i_h.GetAttrInt("color"), i_h.GetAttrInt("points"), false)) && (i_h.GetAttrInt("typeId") != CARD_WILD_DRAW_4) {
					return false, "手中有同花色可出牌,不允许+4."
				}
			}
			break
		}
		break
	case TURN_TYPE_DRAW:
		switch the_type_id {
		case CARD_COMMON, CARD_JUMP, CARD_REV, CARD_WILD:
			return false, "摸牌回合 只允许出+2+4"
			break
		case CARD_DRAW_2:
			for _, card_last_tmp := range cards_last {
				if card_last_tmp.GetTypeId() != CARD_DRAW_2 {
					return false, "上回合+4 不允许+2"
				}
			}
			break
		case CARD_WILD_DRAW_4:
			break
		}
		break
	case TURN_TYPE_JUMP:
		switch the_type_id {
		case CARD_COMMON, CARD_JUMP, CARD_REV, CARD_WILD, CARD_DRAW_2, CARD_WILD_DRAW_4:
			//重置回合类型
			this.turn_type = TURN_TYPE_COMMON
			return false, "该回合被跳过"
			break
		}
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

func (this *UnoContext) Match(color, points int, color_only bool) bool {
	if this.points_tmp == POINTS_NO_MATCH {
		color_only = true
	}
	if color_only {
		if this.color_tmp == COLOR_NO_MATCH || this.color_tmp == color {
			return true
		}
	} else {
		if this.color_tmp == COLOR_NO_MATCH || this.color_tmp == color || this.points_tmp == points {
			return true
		}
	}
	return false
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
		this.GetNowUser().SetAttr("isUno", true)
		return true
	} else {
		this.GetNowUser().SetAttr("isUno", false)
		return false
	}
}
