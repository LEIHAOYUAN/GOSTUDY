package main

import "fmt"

/*
1. select有监听channel通信的作用，当有通信发生时就触发相应的代码块
2. select可以管理和编排多个channel，是并发编程当中非常重要的应用
3. 但是一定要注意使用select时的死锁问题，在开发过程中要格外小心。
*/
func test() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		//time.Sleep(3*time.Second)
		ch1 <- 1
	}()
	go func() {
		//time.Sleep(5*time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("can read from ch1")
	case <-ch2:
		fmt.Println("can read from ch2")
	default:
		fmt.Println("default...")
	}
}
