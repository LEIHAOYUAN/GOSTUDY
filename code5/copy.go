package main

import "fmt"

/**
copy函数：切片之间的元素赋值
【注意】：copy的参数必须是slice，不能是数组
如果数组需要使用copy，需要使用 a[:]或者a[i:j]的切片形式
*/
func main() {
	a1 := []int{1, 1, 1, 1, 1}
	b1 := []int{-1, -1, -1}
	// 将b1复制到a1
	copy(a1, b1)
	fmt.Println("a1:", a1)
	fmt.Println("b1:", b1)

	a2 := []int{2, 2, 2, 2, 2}
	b2 := []int{-2, -2, -2}
	// 将a2复制到b2
	copy(b2, a2)
	fmt.Println("a2:", a2)
	fmt.Println("b2", b2)
}
