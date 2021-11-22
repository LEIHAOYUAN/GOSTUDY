package main

import (
	"fmt"
	"math"
)

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
