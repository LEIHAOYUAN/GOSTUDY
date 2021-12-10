package code24

import (
	"fmt"
	"testing"
)

/**
对于普通类型的指针类型传参
仍然是指针地址的拷贝
*/
func TestInt(t *testing.T) {
	i := 10
	ip := &i
	fmt.Printf("原始指针的内存地址是：%p\n", &ip)
	modifyInt(ip)
	fmt.Println("int值被修改了，新值为:", i)
}

func modifyInt(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	*ip = 1
}

func TestMap(t *testing.T) {
	persons := make(map[string]int)
	persons["张三"] = 19

	mp := &persons

	fmt.Printf("原始map的内存地址是：%p\n", mp)
	modifyMap(persons)
	fmt.Println("map值被修改了，新值为:", persons)
}

func modifyMap(p map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &p)
	p["张三"] = 20
}
