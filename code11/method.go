package main

import "fmt"

/*
用户可以给任何自定义类型定义方法，定义一个或者多个均可。
不过要求自定义类型和对应的方法在同一个包中
【注意】给Go语言内置的类型（比如int）定义方法是不可以的，因为要求类型定义和其对应的方法必须在同一个包内。
*/
// 定义一个矩形
type Rectangle struct{ w, h float64 }

func main() {
	rec := Rectangle{w: 2, h: 3}
	fmt.Println(area(rec.w, rec.h))
	fmt.Println(rec.area())
}

// 定义一个求矩形面积的函数
func area(w, h float64) float64 {
	return w * h
}

// 定义一个矩形类型的方法
func (r Rectangle) area() float64 {
	return r.w * r.h
}
