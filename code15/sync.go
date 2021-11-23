package main

import (
	"fmt"
	"sync"
)

/*
【注意】：
当sync.WaitGroup作为参数传递到函数内且调用Done方法的时候，一定不要忘记函数的值传递特性，这里应该传递sync.WaitGroup变量的地址，
如果传递值则会复制另一个sync.WaitGroup类型变量出来，和函数外的那个变量就没有关系了，
即便调用了Done方法，最后还是会报DeadLock错误。
*/
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		// 给计数器+1
		wg.Add(1)
		go func(x int) {
			// 给计数器-1
			defer wg.Done()
			fmt.Print(" ", x)
		}(i)
	}
	fmt.Printf("\n%#v\n", wg)
	wg.Wait()
	fmt.Println("\nThe End!")
}
