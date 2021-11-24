package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

/*
【性能分析要点】
1. CPU占有率高，高负荷运转，要找到是执行哪个函数时出现这种情况的
2. goroutine在死锁状态，没有运行，但是占用了资源
3. 垃圾回收占用时间

使用runtime/pprof进行性能分析

CPU测试：
对CPU的测试主要是用runcputest函数不停地向通道ch4cpu写入uint64类型的数据，
并由procmsg()函数不停地读数据，而procmsg函数结束的条件是等到chTimer可读，chTimer是启动了一个goroutine并在开始运行15秒后往chTimer写入数据的过程，也就是说，大概在chTimer执行以后，procmsg函数就有机会结束了。
因为有大量的通道读写，所以这是测试CPU的主要代码

内存测试：
对于内存的测试，则是通过大量分配内存堆实现的，并且每次都有一个变量是可回收的，如此反复，每次都执行垃圾回收，并记录堆的信息，这是在模拟一种内存使用情况，可以通过分析记录查看内存泄露等情况

*/

var ch4cpu chan uint64
var chTimer chan struct{}
var memMap map[int]interface{}

func init() {
	ch4cpu = make(chan uint64, 10000)
	chTimer = make(chan struct{}, 20)
	memMap = make(map[int]interface{})
}
func main() {
	c, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/cpu_profile.prof")
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	m1, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/mem_profile1.prof")
	if err != nil {
		log.Fatal(err)
	}

	m2, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/mem_profile2.prof")
	if err != nil {
		log.Fatal(err)
	}

	m3, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/mem_profile3.prof")
	if err != nil {
		log.Fatal(err)
	}

	m4, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/mem_profile4.prof")
	if err != nil {
		log.Fatal(err)
	}

	defer m1.Close()
	defer m2.Close()
	defer m3.Close()
	defer m4.Close()

	pprof.StartCPUProfile(c)
	defer pprof.StopCPUProfile()

	memMap[1] = runMEMTest()

	runtime.GC()
	pprof.Lookup("heap").WriteTo(m1, 0)
	//从此处开始ch4cpu通道被不断地写入数据
	go runcputest()
	//goroutine运行15秒后chTimer写入值
	go func() {
		time.Sleep(15 * time.Second)
		log.Println("write timer")
		chTimer <- struct{}{}

	}()
	memMap[2] = runMEMTest()
	runtime.GC()
	pprof.Lookup("heap").WriteTo(m2, 0)

	memMap[2] = nil
	runtime.GC()
	pprof.Lookup("heap").WriteTo(m3, 0)

	memMap[1] = nil
	runtime.GC()
	pprof.Lookup("heap").WriteTo(m4, 0)

	procmsg()
}

func runMEMTest() []int {
	mem := make([]int, 100000, 120000)
	return mem
}

func runcputest() {
	var i uint64
	for {
		ch4cpu <- i
		i++
	}
}

func procmsg() {
	for {
		select {
		case _ = <-ch4cpu:
		case _ = <-chTimer: //直到满足此条件for循环才结束
			log.Println("timeout")
			return
		}
	}
}
