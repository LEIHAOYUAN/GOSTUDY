package main

import "fmt"

/*
map底层是一个Hash表，对键有唯一性要求！！！
注意点：
1. 在同一个map中所有的k必须是同一类型，而且只有可以比较，或者说只有可以使用“==”符号比较的类型才可以作为k。显然用bool类型作k并不灵活，而使用浮点型作为k，可能会因为不同机器和系统对于精度定义的不同而导致异常
2. 在同一个map中，v也只能是同一类型。在定义v的时候，可以选择任何类型，没有限制
*/
func main() {
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
