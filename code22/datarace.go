package code22

/*
Go语言的多线程模式提倡：不要通过共享内存进行通信，而是要通过通信来共享内存
*/
// 设置数字值
var realNum = make(chan int)

// 设置的增减额
var delta = make(chan int)

// SetNumber 将参数n写入通道realNum
func SetNumber(n int) {
	realNum <- n
}

// ChangeByDelta 向delta通道写入参数d的值
func ChangeByDelta(d int) {
	delta <- d
}

// GetNumber 从realNum通道获取值
func GetNumber() int {
	return <-realNum
}

func monitor() {
	var i int //把数值限定在方法内，goroutine运行后仅在goroutine内可见
	for {
		select {
		case i = <-realNum:
		case d := <-delta:
			i += d
		case realNum <- i:
		}
	}
}

func init() {
	go monitor()
}
