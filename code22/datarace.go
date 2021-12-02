package code22

import "fmt"

/*
数据竞态：
以下示例代码存在数据竞态，因为方法最终返回值取决于 goroutine的执行时机
*/
func test() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()

	return i
}
