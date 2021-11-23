package main

import (
	"fmt"
	"time"
)

/*
1. channel的默认初始值是nil
2. 通道（channel）是Go语言提供的一种在goroutine之间进行数据传输的通信机制
3. 通过channel传递的数据只能是一些指定的类型，这些类型称为通道的元素类型
4. 要使通道正常运行还需要保证通道有数据接收方
*/

func main() {
	c := make(chan int)
	go writeChan(c, 666)
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	if _, ok := <-c; ok {
		fmt.Println("Channel is Open")
	} else {
		fmt.Println("Channel is closed")
	}
}

/*
通过close函数关闭channel不是必须的，只有在需要通知其他的数据读取通道数据已经写入完成时，才必须关闭。
不主动关闭的通道，垃圾回收器会自动回收。
而文件操作的close函数则是必须有的，请读者区分开。
*/
func writeChan(c chan int, x int) {
	fmt.Println("Start:", x)
	c <- x
	close(c)
	fmt.Println("End:", x)
}
