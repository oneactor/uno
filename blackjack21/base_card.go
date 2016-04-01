package blackjack21

import (
	"fmt"
	"github.com/jesusslim/uno"
)

const (
	//牌大类
	JACK_COMMON = 1
	JACK_10     = 2
	JACK_ACE    = 3

	//颜色
	HEARTS   = 1 //红桃
	CLUBS    = 2 //梅花
	DIAMONDS = 3 //方块
	SPADES   = 4 //黑桃

	POINTS_1  = 1
	POINTS_2  = 2
	POINTS_3  = 3
	POINTS_4  = 4
	POINTS_5  = 5
	POINTS_6  = 6
	POINTS_7  = 7
	POINTS_8  = 8
	POINTS_9  = 9
	POINTS_10 = 10
	POINTS_11 = 11
	POINTS_12 = 12
	POINTS_13 = 13
)

type BlackJack21Card struct {
	uno.BaseCard
}

//颜色转输出
func ConvertColor(color int) string {
	result := ""
	switch color {
	case HEARTS:
		result += "红桃"
		break
	case CLUBS:
		result += "梅花"
		break
	case DIAMONDS:
		result += "方块"
		break
	case SPADES:
		result += "黑桃"
		break
	}
	return result
}

func ConvertTitle(color, points int) string {
	result := ""
	result += "[" + ConvertColor(color) + "]"
	switch points {
	case POINTS_1:
		result += "ACE"
		break
	case POINTS_11:
		result += "J"
		break
	case POINTS_12:
		result += "Q"
		break
	case POINTS_13:
		result += "K"
		break
	default:
		result += fmt.Sprintf("%d", points)
	}
	return result
}

func NewBlackJack21Card(id, points, color int) *BlackJack21Card {
	var type_id int
	if points == POINTS_1 {
		type_id = JACK_ACE
	} else if points >= POINTS_10 {
		type_id = JACK_10
	} else {
		type_id = JACK_COMMON
	}
	return &BlackJack21Card{
		*uno.NewBaseCard(id, type_id, ConvertTitle(color, points), color, points),
	}
}

func NewCard(id, points, color int) uno.Card {
	return NewBlackJack21Card(id, points, color)
}

func NewCards(id_from int) []uno.Card {
	cards := []uno.Card{}
	for _, color := range []int{HEARTS, CLUBS, DIAMONDS, SPADES} {
		for _, point := range []int{POINTS_1, POINTS_2, POINTS_3, POINTS_4, POINTS_5, POINTS_6, POINTS_7, POINTS_8, POINTS_9, POINTS_10, POINTS_11, POINTS_12, POINTS_13} {
			cards = append(cards, NewCard(id_from, point, color))
			id_from++
		}
	}
	return cards
}
