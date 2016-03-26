package cards

import (
	"github.com/jesusslim/uno"
)

type UnoCardCommon struct {
	UnoCard
}

func NewCardCommon(id, color, points int) uno.Card {
	return &UnoCardCommon{
		UnoCard{
			id:       id,
			type_id:  CARD_COMMON,
			title:    "",
			color:    color,
			points:   points,
			ext_type: TYPE_COMMON,
		},
	}
}

func (this *UnoCardCommon) CheckPlay() (bool, string) {
	return false, "花色不对"
}

func NewCardCommons(id_from int) []uno.Card {
	var result []uno.Card
	id_now := id_from
	for _, color := range []int{COLOR_RED, COLOR_YELLOW, COLOR_BLUE, COLOR_GREEN} {
		for _, point := range []int{POINTS_0, POINTS_1, POINTS_2, POINTS_3, POINTS_4, POINTS_5, POINTS_6, POINTS_7, POINTS_8, POINTS_9} {
			result = append(result, NewCardCommon(id_now, color, point))
			id_now++
			if point != POINTS_0 {
				result = append(result, NewCardCommon(id_now, color, point))
				id_now++
			}
		}
	}
	return result
}
