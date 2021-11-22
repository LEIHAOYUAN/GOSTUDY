package main

import "fmt"

/**
【闭包】要点：https://blog.csdn.net/cbmljs/article/details/103353746
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
	fmt.Println("可变长参数函数测试............................")
	args := []interface{}{1234, "abcd"}
	print(args)
	// 函数调用加省略号的作用：将切片中的值解析出来，然后传递给参数
	print(args...)
}

func double() func() int {
	var r int
	return func() int {
		r++
		return r * 2
	}
}

/*
多返回值函数
*/
func swap(a, b int) (int, int) {
	return b, a
}

/*
变长参数函数的定义
*/
func sum(a int, others ...int) int {
	for _, v := range others {
		a += v
	}
	return a
}

/*
空接口定义可变长参数，意味着可以传递任何类型的参数。

*/
func print(a ...interface{}) {
	fmt.Println(a...)
}
