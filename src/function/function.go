package function

import "fmt"

// Find 判断数组中是否有指定元素
// 传入一个数组和指定元素
func Find (arr []int,num int)(flag bool){
	flag = false
	for i := 0;i < len(arr);i ++{
		if arr[i] == num{
			flag = true
		}
	}
	return flag
}
// Count 记数组中某指定元素的数量
// 传入一个数组和指定元素
func Count (arr []int,num int)(n int) {
	n = 0
	for i := 0;i < len(arr);i ++{
		if arr[i] == num{
			n ++
		}
	}
	return n
}
// Delete 返回删除数组中指定序号的元素的一个切片
// 传入一个数组和待删除元素的序号
func Delete (arr1 []int,i int)(arr2 [] int){
	field1 := arr1[0 : i]
	field2 := arr1[i+1 :]
	arr2 = append(field1,field2...)
	return arr2
}
// SliceCopy 返回一个数组的复制切片
// 传入一个数组
func SliceCopy(s []int) []int {
	var slice = make([]int, len(s))
	copy(slice, s)
	return slice
}

func NextPlayer(id1 int)(id2 int){
	if id1 < 3{
		id2 = id1 + 1
	}else if id1 == 3{
		id2 = 0
	}else{
		fmt.Println("ERROR!")
	}
	return id2
}