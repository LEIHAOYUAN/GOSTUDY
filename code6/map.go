package main

import "fmt"

/*
map底层是一个Hash表，对键有唯一性要求！！！
注意点：
1. 在同一个map中所有的k必须是同一类型，而且只有可以比较，或者说只有可以使用“==”符号比较的类型才可以作为k。
显然用bool类型作k并不灵活，而使用浮点型作为k，可能会因为不同机器和系统对于精度定义的不同而导致异常
2. 在同一个map中，v也只能是同一类型。在定义v的时候，可以选择任何类型，没有限制
----------------------------------------------------------------------------------------------------------------------------------
注意
1. map类型的默认初始值是nil，也就是说未初始化的map是nil。尽管如此，未初始化的map执行删除元素、len操作、range操作或查找元素时都不会报错。但是如果在初始化之前进行元素赋值则会报错。
2. map的遍历顺序是不固定的，不同的机器可能对Hash算法的使用会有所不同，而且从实际应用来看，map的遍历顺序确实体现出无序的特征。如果要对map排序，需要对key进行排序，然后根据安装key的顺序取值来达到map排序的效果。
*/

func main() {
	initMap()
	m1 := make(map[string]int)
	m1["k1"] = 11
	m1["k2"] = 22
	print(m1)
	delete(m1, "k1")
	print(m1)
	delete(m1, "k1")
	print(m1)

	val, ok := m1["k1"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("k1不存在")
	}
}

func print(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
func initMap() {
	// 方式一：make函数创建map
	m1 := make(map[string]int)
	fmt.Println("make方法创建map：", m1)
	// 方式二：map关键字创建并初始化map
	m2 := map[string]int{
		"aaa": 111,
		"bbb": 222,
	}
	fmt.Println("map关键字创建并初始化map：", m2)
}
