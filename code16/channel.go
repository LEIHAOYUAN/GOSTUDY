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
	// 定义通道c
	c := make(chan int)
	go writeChan(c, 666)
	time.Sleep(1 * time.Second)
}

func writeChan(c chan int, x int) {
	fmt.Println(x)
	// 将x的值写入channel内
	c <- x
	// 使用close函数关闭channel，channel关闭后就不可以在通信了
	close(c)
	fmt.Println(x)
}
