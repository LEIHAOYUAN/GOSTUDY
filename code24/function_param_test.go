package code24

import (
	"fmt"
	"testing"
)

/**
对于普通类型的指针类型传参
仍然是指针地址的拷贝
*/
func TestPointParam(t *testing.T) {
	i := 10
	ip := &i
	fmt.Printf("原始指针的内存地址是：%p\n", &ip)
	modify(ip)
	fmt.Println("int值被修改了，新值为:", i)
}

func modify(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	*ip = 1
}
