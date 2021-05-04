package function

var Allcard = [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	31, 32, 33, 34, 41, 42, 43, 31, 32, 33, 34, 41, 42, 43, 31, 32, 33, 34, 41, 42, 43, 31, 32, 33, 34, 41, 42, 43}
// CanTing 判断牌型是否可以听牌
// 返回是否可听、听什么
func CanTing(handCards, showCards []int) (bool) {
	var canTing = false

	// 循环将可能听的牌，带入到手牌，再用胡牌算法检测是否可胡
	for _, t := range GetMaybeTing(handCards, showCards) {
		if CanWin(append([]int{t}, handCards...), showCards) {
			canTing = true
		}
	}
	return canTing
}

func TingList(handCards, showCards []int) ( []int){
	tingCards := make([]int, 0)

	// 循环将可能听的牌，带入到手牌，再用胡牌算法检测是否可胡
	for _, t := range GetMaybeTing(handCards, showCards) {
		if CanWin(append([]int{t}, handCards...), showCards)&&!Find(tingCards,t) {
			tingCards = append(tingCards, t)
		}
	}
	return tingCards
}

// GetMaybeTing 获取哪些牌是可能听的
// 边张只有自身或者上下张的某一张
// 其他的是自身和上下张
// 如果有明牌，且明牌是3张的话，则明牌也可能是胡的
func GetMaybeTing(handCards, showCards []int) []int {
	var maybeCards [] int
	for i :=0;i < len(handCards);i ++{
		if Find(Allcard,handCards[i]-1)&&Find(Allcard,handCards[i]+1){
			maybeCards = append(maybeCards,handCards[i]-1)
			maybeCards = append(maybeCards,handCards[i]+1)
			maybeCards = append(maybeCards,handCards[i])
		}else{
			maybeCards = append(maybeCards,handCards[i])
		}
	}
	if len(showCards) == 3 &&
		showCards[0] == showCards[1] && showCards[1] == showCards[2] &&
		!Find(maybeCards, showCards[0]) {
		maybeCards = append(maybeCards, showCards[0])
	}
	return maybeCards
}


