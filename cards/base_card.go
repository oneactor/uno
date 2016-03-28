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
	COLOR_BLACK  = 1005

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

func ConvertTitle(ext_type, type_id, color, points int) string {
	result := ""
	switch ext_type {
	case TYPE_COMMON:
		result += "[普通牌]"
		break
	case TYPE_USEAGE:
		result += "[功能牌]"
		break
	case TYPE_ALL_CAN:
		result += "[万能牌]"
		break
	}
	switch type_id {
	case CARD_COMMON:
		result += " 普通 "
		break
	case CARD_JUMP:
		result += " 跳过 "
		break
	case CARD_REV:
		result += " 反向 "
		break
	case CARD_WILD:
		result += " 万能 "
		break
	case CARD_DRAW_2:
		result += " +2 "
		break
	case CARD_WILD_DRAW_4:
		result += " 万能+4 "
		break
	}
	switch color {
	case COLOR_BLACK:
		result += "<黑色>"
		break
	case COLOR_BLUE:
		result += "<蓝色>"
		break
	case COLOR_GREEN:
		result += "<绿色>"
		break
	case COLOR_RED:
		result += "<红色>"
		break
	case COLOR_YELLOW:
		result += "<黄色>"
		break
	}
	if ext_type == TYPE_COMMON {
		result += fmt.Sprintf("%d", points) + "点"
	}
	return result
}

func (this *UnoCard) GetId() int {
	return this.id
}

func (this *UnoCard) GetTypeId() int {
	return this.type_id
}

func (this *UnoCard) GetTitle() string {
	return this.title
}

func (this *UnoCard) GetColor() int {
	return this.color
}

func (this *UnoCard) GetPoints() int {
	return this.points
}

func (this *UnoCard) GetExtType() int {
	return this.ext_type
}

func (this *UnoCard) OnDraw() {
	fmt.Println("on draw")
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

func (this *UnoCard) GetAttr(name string) interface{} {
	switch name {
	case "id":
		return this.id
		break
	case "type_id":
	case "typeId":
		return this.type_id
		break
	case "title":
		return this.title
		break
	case "color":
		return this.color
		break
	case "points":
		return this.points
		break
	case "ext_type":
	case "extType":
		return this.ext_type
		break
	}
	return nil
}

func (this *UnoCard) GetAttrInt(name string) int {
	switch name {
	case "id":
		return this.id
		break
	case "type_id":
	case "typeId":
		return this.type_id
		break
	case "color":
		return this.color
		break
	case "points":
		return this.points
		break
	case "ext_type":
	case "extType":
		return this.ext_type
		break
	}
	return 0
}

func (this *UnoCard) GetAttrStr(name string) string {
	switch name {
	case "title":
		return this.title
	}
	return ""
}
