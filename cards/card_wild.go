package cards

import (
	"github.com/jesusslim/uno"
)

/**
 * uno万能牌
 */
type UnoCardWild struct {
	UnoCard
}

func NewCardWild(id, type_id int) uno.Card {
	return &UnoCardCommon{
		*NewUnoCard(id,
			type_id,
			ConvertTitle(TYPE_ALL_CAN, type_id, COLOR_BLACK, -1),
			COLOR_BLACK,
			-1,
			TYPE_ALL_CAN),
	}
}

func NewCardWilds(id_from int) []uno.Card {
	var result []uno.Card
	id_now := id_from
	for _, type_id := range []int{CARD_WILD, CARD_WILD_DRAW_4} {
		for i := 0; i < 4; i++ {
			result = append(result, NewCardWild(id_now, type_id))
			id_now++
		}
	}
	return result
}
