package main

import (
	"fmt"
	"github.com/jesusslim/uno"
	"github.com/jesusslim/uno/cards"
)

func main() {
	fmt.Println("Example For Uno.")
	desk := cards.NewDesk()
	cards_common := cards.NewCardCommons(1)
	desk.PrepareCards(cards_common)
	cards_usegae := cards.NewCardUseages(101)
	desk.PrepareCards(cards_usegae)
	cards_wild := cards.NewCardWilds(201)
	desk.PrepareCards(cards_wild)
	for k, v := range desk.GetCards() {
		fmt.Println(k, v)
	}
	desk.Shuffle()
	fmt.Println("shuffled")
	user1 := cards.NewUnoUser(10017, "T1")
	user2 := cards.NewUnoUser(10018, "T2")
	user3 := cards.NewUnoUser(10019, "T3")
	user4 := cards.NewUnoUser(10020, "T4")
	ct := cards.NewUnoContext(desk)
	ct.AddUser(user1)
	ct.AddUser(user2)
	ct.AddUser(user3)
	ct.AddUser(user4)
	for _, u := range ct.GetUsers() {
		for i := 0; i < 7; i++ {
			card := desk.GetNext()
			u.AddCard(card)
		}
	}
	i := 0
	for {
		i++
		ct.OnStart()
		u := ct.GetNowUser()
		if u.GetId() == 10017 {
			fmt.Println("====================")
		}
		fmt.Println("ROUND ", i)
		fmt.Println("Now it's ", u.GetNick(), "'s turn.")
		fmt.Println("He has cards:")
		for _, c := range u.GetCards() {
			fmt.Println("  ", c)
		}
		fmt.Println("He use card:")
		var card_id_use int
		var card uno.Card
		if u.GetId() == 10017 {
			fmt.Scan(&card_id_use)
			fmt.Println("U choose id ", card_id_use)
			cards := u.GetCards()
			card = cards[card_id_use]
		} else {
			for id, c := range u.GetCards() {
				card_id_use = id
				card = c
				break
			}
		}
		fmt.Println("  ", card)
		ok, info := ct.CheckPlay([]int{card_id_use})
		fmt.Println(ok, info)
		if ok {
			fmt.Println("On Play...")
			ct.OnPlay([]int{card_id_use}, map[string]interface{}{})
		}
		//Check win
		win := ct.CheckWinner()
		if win {
			fmt.Println("WIN!")
			break
		}
		//Check uno
		r := ct.CheckUno()
		if r {
			fmt.Println("Uno.")
		}
		fmt.Println("On End...")
		not_end, info := ct.OnEnd()
		if !not_end {
			fmt.Println("END!!", info)
			break
		}
	}
}
