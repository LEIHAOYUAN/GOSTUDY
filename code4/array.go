package main

import "fmt"

func main() {
	// 定义一个长度为3的整型数组，里面三个元素的默认值均为0
	var a [3]int
	// 定义数组的同时进行初始化的赋值
	var b [3]int = [3]int{1, 2, 3}
	// 精简化的声明方式
	var b1 = [3]int{1, 2, 3}
	// 省略数组长度
	c := [...]int{1, 2, 3}
	// 数组的长度根据出现的最大下标确定
	d := [...]int{4, 4: 1, 1: 2}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b1)
	fmt.Println(c)
	fmt.Println(d)
	// 数组的遍历
	for i := range d {
		fmt.Println(i)
	}
	/***
	go语言中数组的长度是数组的一部分。
	在分配底层空间的时候，数组的元素是紧挨着分配到固定位置的，这也是数组长度不可变的原因。
	*/
}
