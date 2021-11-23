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

【注意】
缓存通道的发送操作是在队列的尾部插入元素，而接收操作是从队列的头部移除一个元素。
如果队列满了，发送数据的goroutine则会进入阻塞状态等待另一个goroutine来读取数据，进而腾出空间；
而如果队列是空的，接收数据的goroutine则进入阻塞状态等待另一个goroutine在通道上发送数据
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

/*
双向通道
*/
func one(c chan int, x int) {
	fmt.Println(x)
	c <- x
}

/*
带有方向的通道
*/
func two(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
}
