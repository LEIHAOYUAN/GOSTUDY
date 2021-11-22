package main

import "fmt"

func main() {
	f := double()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func double() func() int {
	var r int
	return func() int {
		r++
		return r * 2
	}
}

// 具名函数
func Square(a int) int {
	return a * a
}

// 匿名函数
var square = func(a int) int {
	return a * a
}
