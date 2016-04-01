package cards

import (
	"fmt"
	"github.com/jesusslim/uno"
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
	uno.BaseContext

	this_turn_played bool // 该回合是否成功出了牌

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
		*uno.NewBaseContext(desk),

		false,

		COLOR_NO_MATCH,
		POINTS_NO_MATCH,
		0,
		TURN_TYPE_COMMON,
	}
}

func NewContext(desk uno.Desk) uno.Context {
	return NewUnoContext(desk)
}

func (this *UnoContext) TurnNext() {
	if this.this_turn_played {
		this.SetCardsLast(this.GetCardsNow())
		this.SetCardsNow(map[int]uno.Card{})
	}
	this.this_turn_played = false
	this.BaseContext.TurnNext()
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
					index := this.GetRandom(len(colors))
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
