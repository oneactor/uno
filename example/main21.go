package main

import (
	"fmt"
	//"github.com/jesusslim/uno"
	"github.com/jesusslim/uno/blackjack21"
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

	user1 := blackjack21.NewUser(10017, "T1")
	user2 := blackjack21.NewUser(10018, "T2")
	user3 := blackjack21.NewUser(10019, "T3")
	user4 := blackjack21.NewUser(10020, "T4")
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

	for i := 0; i < len(ctx.GetUsers()); i++ {
		ctx.TurnNext()
		ctx.OnStart()
		u := ctx.GetNowUser()
		fmt.Println("====================")
		fmt.Println("ROUND ", i)
		fmt.Println("Now it's ", u.GetNick(), "'s turn.")
		for {
			fmt.Println("-----------")
			fmt.Println("He has cards:")
			for _, c := range u.GetCards() {
				fmt.Println("  ", c)
			}
			actions := blackjack21.GetActions()
			fmt.Println("----> You can choose:")
			fmt.Println(actions)
			var choice int
			fmt.Scan(&choice)
			fmt.Println("U choose : ", actions[choice])
			ctx.OnPlay([]int{}, map[string]interface{}{"action": choice})
			r := ctx.CheckTurnEnd()
			if r {
				break
			}
		}
	}
	ctx.CheckWinner()
}
