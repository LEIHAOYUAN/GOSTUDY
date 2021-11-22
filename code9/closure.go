package main

import "fmt"

/**
【闭包】要点：
0. 闭包本身可以看作是独立对象
1. 闭包函数与普通函数的最大区别就是参数不是值传递，而是引用传递，所以闭包函数可以操作自己函数以外的变量。
2. 因为闭包函数对外部变量的操作才使其不能被释放回收，从而跨过了作用域的限制
总结：闭包会将自己用到的变量都保存在内存中，导致变量无法被及时回收，并且可能通过闭包修改父函数使用的变量值，所以在使用的时候也要注意性能和安全。
*/
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
