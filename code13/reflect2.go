package main

import (
	"fmt"
	"reflect"
)

/*
1. 反射代码的写法可读性比较差，不利于后续的运维。
2. 反射的实现比较复杂，所以反射执行得比较慢，会影响程序的整体性能。
3. 反射的错误在编译时无法发现，到运行时才报错，而且都是panic类型，这容易让程序崩溃。
*/
type XX struct {
	I int
	F float64
	S string
}
type Person struct {
	Name   string `json:"jname"`
	Gender int    `json:"jgender"`
	Age    int    `json:"jage"`
}

/*
比较两个对象是否相等
*/
func (x XX) CompareStr(xx XX) bool {
	rx1 := reflect.ValueOf(&x).Elem()
	rx2 := reflect.ValueOf(&xx).Elem()
	for i := 0; i < rx1.NumField(); i++ {
		if rx1.Field(i).Interface() != rx2.Field(i).Interface() {
			return false
		}
	}
	return true
}

func (p Person) PrintTags() {
	for i := 0; i < reflect.TypeOf(p).NumField(); i++ {
		fmt.Println(reflect.TypeOf(p).Field(i).Tag.Get("json"))
	}
}

func main() {
	x1 := XX{I: 1, F: 1.2, S: "hello"}
	x2 := XX{I: 1, F: 1.2, S: "hello"}
	fmt.Println(x1.CompareStr(x2))

	p := Person{Name: "Scott", Gender: 1, Age: 30}
	p.PrintTags()

}
