package main

import "fmt"

//定义一个矩形
type Rectangle1 struct{ w, h float64 }

func main() {
	p := &Rectangle1{w: 2, h: 3}
	fmt.Println(p.area())

	rec := Rectangle1{w: 2, h: 3}
	rp := &rec
	fmt.Println(rp.area())

	fmt.Println((&rec).area())
	fmt.Println(rec.area()) //会隐式地加上*rec，合法用法

	// Rectangle1{w: 2, h: 3}.area() //会报错，因为无法通过这种方式获取变量的地址
	Rectangle1{w: 2, h: 3}.area2() //合法用法，因为area2方法的接收器不是指针类型，采用值传递

}

// 定义一个接收器为指针类型的方法
func (r *Rectangle1) area() float64 {
	return r.w * r.h
}

// 定义接收器为普通对象的方法：所有变量的值都会通过值传递的方式传递到方法内
func (r Rectangle1) area2() float64 {
	return r.w * r.h
}
