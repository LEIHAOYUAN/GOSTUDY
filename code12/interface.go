package main

import (
	"fmt"
	"math"
)

/*
接口中Go的设计理念：
go语言的理念是没必要把数据和方法封装在一起，只需要通过接口将逻辑进行抽象和组合即可。
接口是Go语言实现面向对象的最重要方式。Go语言的面向对象依赖于接口，而不是依赖于实现

Go语言的类型断言可以用x.(T)表达，其中x是一个接口类型的具体值表达式，而T是一个类型—断言类型。T的主要作用就是检查动态值x是否满足T。
注意：
接口类型的值（或简称接口值）包括动态类型和动态值，也就是说在编译阶段并不知道具体的类型和值，而是在程序执行到此时再通过动态类型和动态值去调用具体的方法。所以读者在思考接口运行方式时，始终要将接口看作动态类型和动态值两个字段，这样更有利于理解。
1. 接口中只能声明方法，不可以有具体实现。
2. 接口中不可以声明变量，仅允许声明方法。
3. 实现一个接口，就必须实现接口内声明的所有方法。
4. 实现一个方法就是要和接口声明的方法的方法名、形参、返回值完全一致。
5. 接口也是可以嵌套组合的，和结构体一样。
6. 在接口中声明方法时，不可以出现重名方法。
*/
type ShapeDesc1 interface {
	Area() float64
	Perimeter() float64
}

type rectangle struct {
	H, W float64
}

type circle struct {
	R float64
}

func (r rectangle) Area() float64 {
	return r.H * r.W
}
func (r rectangle) Perimeter() float64 {
	return 2 * (r.H + r.W)
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}
func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

func main() {
	var s1, s2 ShapeDesc1
	s1 = rectangle{H: 2, W: 3} //注意此处，rectangle实现了ShapeDesc1接口
	s2 = circle{R: 2}          //注意此处，circle实现了ShapeDesc1接口
	Desc1(s1)
	Desc1(s2)
}

/*
类型判断：方式一
*/
func Desc1(s ShapeDesc1) {
	_, ok := s.(circle)
	if ok {
		fmt.Println("This is circle.")
	}
	_, ok = s.(rectangle)
	if ok {
		fmt.Println("This is rectangle.")
	}
	fmt.Println("area:", s.Area())
	fmt.Println("perimeter:", s.Perimeter())
}

/*
类型判断：方式二
*/
func Desc2(s ShapeDesc1) {
	switch kind := s.(type) {
	case circle:
		fmt.Println("This is circle.")
	case rectangle:
		fmt.Println("This is rectangle.")
	default:
		fmt.Println("%v is unknown type", kind)
	}

	fmt.Println("area:", s.Area())
	fmt.Println("perimeter:", s.Perimeter())
}
