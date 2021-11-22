package main

import "fmt"

/*
如果函数内部改变了slice的长度或者容量，则原slice将不会发生改变
*/
func add1(arr []int) []int {
	// 直接修改切片的值可以，实参也会跟着修改
	// arr[0] = 99
	arr = append(arr, 999)
	fmt.Println("修改后arr：", arr)
	return arr
}

func add2(arr []int) {
	// 直接修改切片的值可以，实参也会跟着修改
	arr[0] = 99
	fmt.Println("修改后arr：", arr)
}

func main() {
	var a = make([]int, 2, 8)
	a = add1(a)
	fmt.Println("修改后a：", a)
}
