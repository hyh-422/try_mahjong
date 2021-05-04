package player

import (
	"fmt"

	"../card"
	"../function"
)



type Player struct {
	Name string
	Id string
	Score int
	Hand card.Handfield
	River card.Riverfield
	Display card.Displayfield
}

func (player *Player) Init(){
	player.Hand.Init()
	player.River.Init()
	player.Display.Init()
}

func (player *Player) Get(i int){
	player.Hand.Push(i)
}

func (player *Player) Pop() (i int){
	fmt.Println("请输入牌号")
	fmt.Scanf("%d", &i)
	fmt.Printf("%s弃%d\n",player.Name,i)
	player.Hand.Remove(i)
	player.River.Push(i)
	return i
}

func (player *Player) Eat(num int)(flag bool){
	flag = false
	if num >= 30{
		fmt.Println("不能吃")
		return flag
	}
	if function.Find(player.Hand.Field,num-1)&&function.Find(player.Hand.Field,num+1){
		var temp = [] int {num-1,num,num+1}
		player.Hand.Remove(num-1)
		player.Hand.Remove(num+1)
		player.Display.Push(temp)
		flag = true
	}else if function.Find(player.Hand.Field,num-2)&&function.Find(player.Hand.Field,num-1){
		var temp = [] int {num-2,num-1,num}
		player.Hand.Remove(num-2)
		player.Hand.Remove(num-1)
		player.Display.Push(temp)
		flag = true
	}else if function.Find(player.Hand.Field,num+1)&&function.Find(player.Hand.Field,num+2){
		var temp = [] int {num,num+1,num+2}
		player.Hand.Remove(num+1)
		player.Hand.Remove(num+2)
		player.Display.Push(temp)
		flag = true
	}
	if flag == false{
		fmt.Println("不能吃")
	}
	return flag
}

func (player *Player) Peng(num int)(flag bool){
	flag = false
	if function.Count(player.Hand.Field,num) >= 2{
		var temp = [] int {num,num,num}
		player.Hand.Remove(num)
		player.Hand.Remove(num)
		player.Display.Push(temp)
	}
	return flag
}


