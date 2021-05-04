package main

import (
	"fmt"

	"../function"
)

func start(){
	cardgame.Init()
	cardgame.Fapai()
	flag = 1
}

func stateone(){
	cardgame.Mopai(cardgame.NowplayerId)
	watch()
	cardgame.Chupai(cardgame.NowplayerId)
	flag = 2
}

func statetwo(){
	var tmp = function.NextPlayer(cardgame.NowplayerId)
	var tmparr1 [] int
	var tmparr2 [] int
	for i := 0;i < 3;i ++{
		tmparr1 = append(cardgame.Players[tmp].Hand.Field,cardgame.Wait)
		tmparr2 = cardgame.Players[tmp].Display.Field
		if function.CanWin(tmparr1,tmparr2){
			fmt.Println("是否胡牌")
			//if true
			flag = 5
			//add winner
			break
		}
		tmp = function.NextPlayer(tmp)
	}
	flag = 3
}

func statethree(){
	var tmp = function.NextPlayer(cardgame.NowplayerId)
	var wait = cardgame.Wait
	var tmparr1 [] int
	for i := 0;i < 3;i ++{
		tmparr1 = cardgame.Players[tmp].Hand.Field
		if function.Count(tmparr1,wait) >= 2{
			fmt.Println("是否碰")
			//if true
			cardgame.Players[tmp].Peng(wait)
			watch()
			cardgame.Chupai(tmp)
			cardgame.NowplayerId = tmp
			flag = 2
			//else continue
		}
		tmp = function.NextPlayer(tmp)
	}
	flag = 4
}

func statefour(){
	var tmp = function.NextPlayer(cardgame.NowplayerId)
	var wait = cardgame.Wait
	var tmparr = cardgame.Players[tmp].Hand.Field
	if (function.Find(tmparr,wait-2)&&function.Find(tmparr,wait-1))||
		(function.Find(tmparr,wait-1)&&function.Find(tmparr,wait+1))||
		(function.Find(tmparr,wait+1)&&function.Find(tmparr,wait+2)){
		fmt.Println("是否吃")
		//if true
		cardgame.Players[tmp].Eat(wait)
		watch()
		cardgame.Chupai(tmp)
		cardgame.NowplayerId = tmp
		flag = 2
		//else continue
	}
	cardgame.NowplayerId = tmp
	flag = 1
}

func watchall(){
	for i := 0;i < 4;i ++{
		fmt.Println(cardgame.Players[i].Name)
		fmt.Println(cardgame.Players[i].Hand)
	}
}

func watch(){
	fmt.Println(cardgame.Players[cardgame.NowplayerId].Name)
	fmt.Println(cardgame.Players[cardgame.NowplayerId].Hand)
}
