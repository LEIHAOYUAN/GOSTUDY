package code23

import "testing"

/*
从数组生成一个新的切片
*/
func TestSliceLenCap1(t *testing.T) {
	var arr = [5]int{1, 22, 333, 444, 5555}
	var s1 = arr[1:3]
	t.Log("使用数组生成一个新的切片：", len(s1), cap(s1))
}

/*
从切片生成一个新的切片
*/
func TestSliceLenCap2(t *testing.T) {
	var arr = [5]int{1, 22, 333, 444, 5555}
	var s1 = arr[1:3]
	var s2 = s1[0:1]
	t.Log(s1)
	t.Log(len(s1), cap(s1))
	t.Log(s2, len(s2), cap(s2))
}

/*
直接声明一个新切片
*/
func TestSliceLenCap3(t *testing.T) {
	var s10 = []int{1, 2, 3, 4, 5} //申明的同时初始化
	t.Log(s10, len(s10), cap(s10))

	s20 := make([]int, 2, 4) //通过make申明
	t.Log(s20, len(s20), cap(s20))
}

/*
增删切片对len和cap的影响
*/
func TestSliceLenCap4(t *testing.T) {
	s20 := make([]int, 2, 4) //通过make申明
	t.Log(s20, len(s20), cap(s20))
	s20 = append(s20, 1, 2, 3)
	t.Log(s20, len(s20), cap(s20))

	s30 := make([]int, 2, 4)
	for i := 0; i < 10; i++ {
		s30 = append(s30, i)
		t.Log(s30, len(s30), cap(s30))
	}
}
