package code21

import "testing"

/*
【基准测试】
基准测试就是性能测试。
CPU测试：
对CPU的测试主要是用runcputest函数不停地向通道ch4cpu写入uint64类型的数据，
并由procmsg()函数不停地读数据，而procmsg函数结束的条件是等到chTimer可读，chTimer是启动了一个goroutine并在开始运行15秒后往chTimer写入数据的过程，也就是说，大概在chTimer执行以后，procmsg函数就有机会结束了。
因为有大量的通道读写，所以这是测试CPU的主要代码

内存测试：
对于内存的测试，则是通过大量分配内存堆实现的，并且每次都有一个变量是可回收的，如此反复，每次都执行垃圾回收，并记录堆的信息，这是在模拟一种内存使用情况，可以通过分析记录查看内存泄露等情况
*/
var final int

func benchmarkfb1(b *testing.B, n int) {
	var end int
	for i := 0; i < b.N; i++ {
		end = fb1(n)
	}
	final = end
}

func Benchmarkfb2(b *testing.B, n int) {
	var end int
	for i := 0; i < b.N; i++ {
		end = fb2(n)
	}
	final = end
}

func Benchmarkfb3(b *testing.B, n int) {
	var end int
	for i := 0; i < b.N; i++ {
		end = fb3(n)
	}
	final = end
}

func Benchmark50fb1(b *testing.B) {
	benchmarkfb1(b, 50)
}

func Benchmark50fb2(b *testing.B) {
	Benchmarkfb2(b, 50)
}

func Benchmark50fb3(b *testing.B) {
	Benchmarkfb3(b, 50)
}
