package code10

import "os"

/*
1. defer类似其他语言中的finally
2. 一个函数内也可以有多个defer，在调用的时候按照栈的方式先进后出，即写在前面的会后调用
*/
func testDefer() {
	f, err := os.Open("路径")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
