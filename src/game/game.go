package game

import (
	"../player"

	"fmt"

	"../card"
)

type Game struct {
	Players [4] player.Player
	Mcard card.Mountain
	Wait int
	NowplayerId int
}

func (game *Game)Init(){
	game.Players[0].Name = "a"
	game.Players[1].Name = "b"
	game.Players[2].Name = "c"
	game.Players[3].Name = "d"
	for i :=0;i <len(game.Players);i ++{
		game.Players[i].Init()
	}
	game.Mcard.Init()
	game.Wait = 0
	game.NowplayerId = 0
}

func (game *Game)Fapai(){
	for i := 0;i < len(game.Players);i ++{
		for j := 0;j <13;j ++{
			var temp int
			temp = game.Mcard.Giveone()
			game.Players[i].Hand.Push(temp)
		}
	}
}

func (game *Game)Mopai(id int){
	var temp = game.Mcard.Giveone()
	game.Players[id].Hand.Push(temp)
	fmt.Println(game.Players[id].Name)
	fmt.Println(game.Players[id].Hand)
}

func (game *Game)Chupai(id int){
	var temp = game.Players[id].Pop()
	game.Wait = temp
}

func (game *Game)Nextplayer(){
	if game.NowplayerId < 3{
		game.NowplayerId += 1
	}else if game.NowplayerId == 3{
		game.NowplayerId = 0
	}else{
		fmt.Println("ERROR!")
	}
}

