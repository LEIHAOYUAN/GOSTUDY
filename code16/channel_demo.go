package main

import (
	"fmt"
	"sync"
)

const (
	// goroutine数量
	noGoroutine = 5
	// 任务数量
	noTask = 10
)

/*
用于控制多个goroutine，待全部执行完成后再结束主goroutine
*/
var wg sync.WaitGroup

func main() {
	//创建缓存容量为noTask的缓存通道
	tasks := make(chan int, noTask)

	//启动数量为noGoroutine的goroutine
	for no := 1; no <= noGoroutine; no++ {
		wg.Add(1)
		go taskProcess(tasks, no)
	}

	//向tasks缓存通道内放入任务号
	for taskNO := 1; taskNO <= noTask; taskNO++ {
		tasks <- taskNO
	}

	/*
		【注意】：如果将close方法和wg.Wait方法调换顺序，将会发生死锁
		原因：goroutine阻塞等待缓冲区继续写入数据，而主goroutine又等待所有的goroutine执行完再关闭通道，因此进入死锁状态
	*/
	// 关闭通道。注意：不管goroutine有没有执行完，先关闭通道，因为即便是通道关闭，如果缓存区还有数据的话，goroutine还是可以读取的
	close(tasks)
	// 等所有的goroutine执行完成后，继续执行主goroutine结束程序
	wg.Wait()

}

func taskProcess(tasks chan int, workerNo int) {
	defer wg.Done()

	for t := range tasks {
		fmt.Printf("Worker %d is processing Task no:%d \n", workerNo, t)
	}
	fmt.Printf("Worker %d got off work \n", workerNo)
}
