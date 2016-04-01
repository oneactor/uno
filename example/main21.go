package main

import (
	"fmt"
	//"github.com/jesusslim/uno"
	"github.com/jesusslim/uno/blackjack21"
	"github.com/jesusslim/uno/cards"
)

func main() {
	fmt.Println("Example for BlackJack21")

	desk := blackjack21.NewDesk()
	desk_cards := blackjack21.NewCards(1)
	desk.PrepareCards(desk_cards)
	for k, v := range desk.GetCards() {
		fmt.Println(k, v)
	}
	desk.Shuffle()
	fmt.Println("shuffled")

	user1 := cards.NewUnoUser(10017, "T1")
	user2 := cards.NewUnoUser(10018, "T2")
	user3 := cards.NewUnoUser(10019, "T3")
	user4 := cards.NewUnoUser(10020, "T4")
	ctx := blackjack21.NewBlackJack21Context(desk)
	ctx.AddUser(user1)
	ctx.AddUser(user2)
	ctx.AddUser(user3)
	ctx.AddUser(user4)

	for _, u := range ctx.GetUsers() {
		for i := 0; i < 2; i++ {
			card := desk.GetNext()
			u.AddCard(card)
		}
	}

	i := 0
	for {
		i++
		ctx.OnStart()
		u := ctx.GetNowUser()
		if u.GetId() == 10017 {
			fmt.Println("====================")
		}
		fmt.Println("ROUND ", i)
		fmt.Println("Now it's ", u.GetNick(), "'s turn.")
		fmt.Println("He has cards:")
		for _, c := range u.GetCards() {
			fmt.Println("  ", c)
		}
		if u.GetId() == 10017 {

		} else {

		}
		fmt.Println("On Play...")
		ctx.OnPlay([]int{}, map[string]interface{}{})
		//Check win
		win := ctx.CheckWinner()
		if win {
			fmt.Println("WIN!")
			break
		}
		not_end, info := ctx.OnEnd()
		if !not_end {
			fmt.Println("END!!", info)
			break
		}
	}
}
