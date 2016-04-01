package cards

import (
	"github.com/jesusslim/uno"
)

/**
 * uno功能牌
 */
type UnoCardUseage struct {
	UnoCard
}

func NewCardUseage(id, color, type_id int) uno.Card {
	return &UnoCardCommon{
		*NewUnoCard(id,
			type_id,
			ConvertTitle(TYPE_USEAGE, type_id, color, -1),
			color,
			-1,
			TYPE_USEAGE),
	}
}

func NewCardUseages(id_from int) []uno.Card {
	var result []uno.Card
	id_now := id_from
	for _, color := range []int{COLOR_RED, COLOR_YELLOW, COLOR_BLUE, COLOR_GREEN} {
		for _, type_id := range []int{CARD_JUMP, CARD_REV, CARD_DRAW_2} {
			result = append(result, NewCardUseage(id_now, color, type_id))
			id_now++
			result = append(result, NewCardUseage(id_now, color, type_id))
			id_now++
		}
	}
	return result
}
