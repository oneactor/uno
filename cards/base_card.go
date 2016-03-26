package cards

import (
	"fmt"
	"github.com/jesusslim/uno"
)

const (
	TYPE_COMMON  = 1 //普通
	TYPE_USEAGE  = 2 //功能
	TYPE_ALL_CAN = 3 //万能

	CARD_COMMON      = 101 //普通
	CARD_JUMP        = 111 //跳过
	CARD_DRAW_2      = 112 //摸2
	CARD_REV         = 113 //反转
	CARD_WILD        = 121 //万能
	CARD_WILD_DRAW_4 = 122 //摸4万能

	COLOR_RED    = 1001
	COLOR_YELLOW = 1002
	COLOR_BLUE   = 1003
	COLOR_GREEN  = 1004

	POINTS_0 = 0
	POINTS_1 = 1
	POINTS_2 = 2
	POINTS_3 = 3
	POINTS_4 = 4
	POINTS_5 = 5
	POINTS_6 = 6
	POINTS_7 = 7
	POINTS_8 = 8
	POINTS_9 = 9
)

type UnoCard struct {
	id       int
	type_id  int
	title    string
	color    int
	points   int
	ext_type int
}

func (this *UnoCard) GetId() int {
	return this.id
}

func (this *UnoCard) GetTypeId() int {
	return this.type_id
}

func (this *UnoCard) OnDraw() {
	fmt.Println("on draw")
}

func (this *UnoCard) OnPlay() {
	fmt.Println("on play")
}

func (this *UnoCard) CheckPlay() (bool, string) {
	return true, ""
}

func NewUnoCard(id int, type_id int, title string, color int, points int, ext_type int) *UnoCard {
	return &UnoCard{
		id:       id,
		type_id:  type_id,
		title:    title,
		color:    color,
		points:   points,
		ext_type: ext_type,
	}
}

func NewCard(id int, type_id int, title string, color int, points int, ext_type int) uno.Card {
	return &UnoCard{
		id:       id,
		type_id:  id,
		title:    title,
		color:    color,
		points:   points,
		ext_type: ext_type,
	}
}
