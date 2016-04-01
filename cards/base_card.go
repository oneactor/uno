package cards

import (
	"fmt"
	"github.com/jesusslim/uno"
)

const (
	//uno牌大类
	TYPE_COMMON  = 1 //普通
	TYPE_USEAGE  = 2 //功能
	TYPE_ALL_CAN = 3 //万能

	//uno牌类型
	CARD_COMMON      = 101 //普通
	CARD_JUMP        = 111 //跳过
	CARD_DRAW_2      = 112 //摸2
	CARD_REV         = 113 //反转
	CARD_WILD        = 121 //万能
	CARD_WILD_DRAW_4 = 122 //摸4万能

	//颜色
	COLOR_RED    = 1001
	COLOR_YELLOW = 1002
	COLOR_BLUE   = 1003
	COLOR_GREEN  = 1004
	COLOR_BLACK  = 1005

	//点数
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

	//无须匹配或无匹配时数值
	COLOR_NO_MATCH  = -100
	POINTS_NO_MATCH = -100
)

/**
 * Uno卡牌
 */
type UnoCard struct {
	uno.BaseCard
	ext_type int // 卡牌大类
}

//颜色转输出
func ConvertColor(color int) string {
	result := ""
	switch color {
	case COLOR_BLACK:
		result += "黑色"
		break
	case COLOR_BLUE:
		result += "蓝色"
		break
	case COLOR_GREEN:
		result += "绿色"
		break
	case COLOR_RED:
		result += "红色"
		break
	case COLOR_YELLOW:
		result += "黄色"
		break
	}
	return result
}

//转卡牌名称
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
	result += "<" + ConvertColor(color) + ">"
	if ext_type == TYPE_COMMON {
		result += fmt.Sprintf("%d", points) + "点"
	}
	return result
}

func (this *UnoCard) GetExtType() int {
	return this.ext_type
}

//新建卡牌
func NewUnoCard(id int, type_id int, title string, color int, points int, ext_type int) *UnoCard {
	return &UnoCard{
		*uno.NewBaseCard(id, type_id, title, color, points),
		ext_type,
	}
}

//新建卡牌
func NewCard(id int, type_id int, title string, color int, points int, ext_type int) uno.Card {
	return NewUnoCard(id, type_id, title, color, points, ext_type)
}

//获取卡牌属性
func (this *UnoCard) GetAttr(name string) interface{} {
	switch name {
	case "id", "type_id", "typeId", "title", "color", "points":
		return this.BaseCard.GetAttr(name)
		break
	case "ext_type", "extType":
		return this.ext_type
		break
	}
	return nil
}

func (this *UnoCard) GetAttrInt(name string) int {
	switch name {
	case "id", "type_id", "typeId", "color", "points":
		return this.BaseCard.GetAttrInt(name)
		break
	case "ext_type", "extType":
		return this.ext_type
		break
	}
	return 0
}
