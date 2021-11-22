package main

import "fmt"

/**
切片：slice是一个拥有相同类型元素的可变长序列，且slice的定义与数组的定义非常像，它是没有长度的数组
指针：指针指向slice开始访问的第一个元素
长度：切片的长度
容量：从slice开始访问的第一个元素到底层数组最后一个元素的元素个数

【说明】：切片的底层是数组，go语言的切片对应着底层数组，一个底层数组可以对应多个slice
*/
func test() {
	// 方式一：切片的定义和数组不同的是不能指定长度
	s1 := []int{1, 2, 3, 4, 5}
	// 方式二：如果只是想创建一个slice而不赋值，那么可以使用make函数进行操作。示例如下：
	s2 := make([]int, 10)
	fmt.Println(s1)
	fmt.Println(s2)
	// 切面的遍历
	for i, v := range s2 {
		fmt.Println(i, v)
	}
	// 使用数组生成切片。切片其实是指向数组的引用
	a := [...]int{1, 2, 3, 4, 5}
	ss := a[1:3]
	fmt.Println(a)
	fmt.Println(ss)
	ss[0] = 666
	ss[1] = -666
	fmt.Println(a)
	fmt.Println(ss)
	fmt.Println(len(ss))
	fmt.Println(cap(ss))

}
