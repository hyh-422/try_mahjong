package main

import (
	"../game"
	"fmt"
)

var flag = 0
var cardgame game.Game

func main(){
	start()
	watchall()
	for flag!=5&&cardgame.Mcard.Num>=0{
		fmt.Println(cardgame.Mcard.Num)
		if flag == 1{
			stateone()
		}else if flag == 2{
			statetwo()
		}else if flag == 3{
			statethree()
		}else if flag == 4{
			statefour()
		}
	}
}
