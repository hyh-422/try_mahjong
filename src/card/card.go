package card

import (
	"container/list"
	"math/rand"
	"sort"
	"time"

	"../function"
)

var Allcard = [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	31, 32, 33, 34, 41, 42, 43, 31, 32, 33, 34, 41, 42, 43, 31, 32, 33, 34, 41, 42, 43, 31, 32, 33, 34, 41, 42, 43}

type Mountain struct {
	Card list.List
	Num int
}

func (mountain *Mountain) Init() {
	rand.Seed(time.Now().UnixNano())
	var copycard [] int = Allcard
	for i := 0;i <136;i++ {
		j := rand.Intn(136)
		temp := 0
		temp = copycard[j]
		copycard[j] = copycard[i]
		copycard[i] = temp
	}
	for i := 0;i <136;i++ {
		mountain.Card.PushBack(copycard[i])
	}

	mountain.Num = 136
}


func (mountain *Mountain) Giveone() int {
	var temp interface{}
	temp = mountain.Card.Front().Value
	mountain.Num -= 1
	mountain.Card.Remove(mountain.Card.Front())
	return temp.(int)
}

type Handfield struct {
	Field [] int
}

func (handfield *Handfield) Init(){
	handfield.Field = [] int{}
}

func (handfield *Handfield) Organize(){
	sort.Ints(handfield.Field)
}

func (handfield *Handfield) Push(i int){
	handfield.Field = append(handfield.Field, i)
	handfield.Organize()
}

func (handfield *Handfield) Remove(i int) int{
	handlen := len(handfield.Field)
	var num int
	for j := 0;j < handlen;j ++{
		if handfield.Field[j] == i{
			num = j
			break
		}
	}
	handfield.Field = function.Delete(handfield.Field,num)
	handfield.Organize()
	return i
}

type Riverfield struct {
	Field [] int
}

func (riverfield *Riverfield) Init(){
	riverfield.Field = [] int{}
}

func (riverfield *Riverfield) Push(i int){
	riverfield.Field = append(riverfield.Field, i)
}

type Displayfield struct {
	Field [] int
}

func (displayfield *Displayfield) Init(){
	displayfield.Field = [] int{}
}

func (displayfield *Displayfield) Push(arr [] int){
	displayfield.Field = append(displayfield.Field,arr...)
}

