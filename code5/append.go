package main

import "fmt"

func main() {
	test()
	fmt.Println("append操作.............................")
	var a = []int{1, 2, 3}
	fmt.Println("cap:", cap(a))
	a = append(a, 444)
	fmt.Println("cap:", cap(a))
	a = append(a, []int{-555, -666, -777, -888}...)
	fmt.Println("cap:", cap(a))

	fmt.Println("遍历a开始...............................")
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("遍历a结束...............................")
	fmt.Println("删除a开始...............................")
	//可以使用append进行删除
	a = append(a[:0], a[:3]...) //只保留前三个元素
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("cap:", cap(a))
	fmt.Println("删除a结束...............................")
	a = append([]int{222, 222}, a...)                   //在开头插入新的切片
	a = append(a[:2], append([]int{-222}, a[2:]...)...) //在下标2的位置插入-222
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("cap:", cap(a))
}
