package main

import "fmt"

/*
结构体的组合使用
1. struct组合涉及多个struct，一个struct可以含有其他struct，以此达到复用效果
2. 虽然struct可以含有其他struct，但不可以含有它自身，也就是说一个struct的成员不可以是本struct。不过，struct内的成员可以是指向自己的指针。
3. struct更加灵活的用法：定义的时候使用tags，使得struct字段与JSON和数据库的字段建立对应关系
*/
type Person struct {
	Name        string
	Gender, Age int
}

type Employee struct {
	p      Person
	Salary int
}

type Student struct {
	// 匿名字段的用法，匿名字段的使用让内部结构体的访问更加简洁便利。但是需要注意的是：必要出现字段名称重复的情况，字段名称重复会让Go语言在编译的时候导致错误
	Person
	School string
}

func main() {
	e := Employee{p: Person{"Scott", 1, 30}, Salary: 1000}
	fmt.Println(e)
	fmt.Println(e.p.Name)

	var s Student
	//相当于 s.Person.Name = "Billy"
	s.Name = "Billy"
	//相当于 s.Person.Gender = 1
	s.Gender = 1
	//相当于 s.Person.Age = 6
	s.Age = 6
	s.School = "定慧里小学"
	fmt.Println(s)
}
