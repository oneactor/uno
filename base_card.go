package uno

import (
	"fmt"
)

type BaseCard struct {
	id      int    // id
	type_id int    // 卡牌类型
	title   string // 名称
	color   int    // 颜色
	points  int    // 点数
}

func (this *BaseCard) GetId() int {
	return this.id
}

func (this *BaseCard) GetTypeId() int {
	return this.type_id
}

func (this *BaseCard) GetTitle() string {
	return this.title
}

func (this *BaseCard) GetColor() int {
	return this.color
}

func (this *BaseCard) GetPoints() int {
	return this.points
}

//摸牌触发
func (this *BaseCard) OnDraw() {
	fmt.Println("on draw")
}

func NewBaseCard(id int, type_id int, title string, color int, points int) *BaseCard {
	return &BaseCard{
		id:      id,
		type_id: type_id,
		title:   title,
		color:   color,
		points:  points,
	}
}

//新建卡牌
func NewBaseCardInterface(id int, type_id int, title string, color int, points int) Card {
	return NewBaseCard(id, type_id, title, color, points)
}

//获取卡牌属性
func (this *BaseCard) GetAttr(name string) interface{} {
	switch name {
	case "id":
		return this.id
		break
	case "type_id", "typeId":
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
	}
	return nil
}

func (this *BaseCard) GetAttrInt(name string) int {
	switch name {
	case "id":
		return this.id
		break
	case "type_id", "typeId":
		return this.type_id
		break
	case "color":
		return this.color
		break
	case "points":
		return this.points
		break
	}
	return 0
}

func (this *BaseCard) GetAttrStr(name string) string {
	switch name {
	case "title":
		return this.title
	}
	return ""
}
