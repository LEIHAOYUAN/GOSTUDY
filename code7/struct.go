package main

import "fmt"

/*
 【注意】
1. struct也是复合类型，而非引用类型。复合类型和引用类型是有区别的，复合类型是值传递，而引用类型是引用传递。
2. struct成员的可见性也是通过首字母大小写控制的，首字母小写仅本包可见，首字母大写则包外也可访问。
3. 结构体指针必须初始化以后才可以使用，因为如果仅仅声明结构体指针类型变量，其默认初始值是nil
-----------------------------------------------------------------------------------------------
make和new关键字注意点：
1. make函数用于slice、map和chan进行内存分配，它返回的不是指针，而是上面三个类型中的某一个类型本身
2. new函数返回初始化的类型对应的指针，new函数主要用在struct初始化中，其他场景应用较少
*/

type Person struct {
	Name   string
	Gender int
	Age    int
}

func main() {
	p1 := Person{Name: "Scott", Gender: 1, Age: 30}
	p2 := AddAge(p1)
	fmt.Println(p1)
	fmt.Println(p2)
	// 注意参数
	AddAgePlus(&p1)
	fmt.Println(p1)

	pp := new(Person)
	AddAgePlus(pp)
	fmt.Println(pp)
}

func AddAge(p Person) (p2 Person) {
	p.Age += 1
	return p
}
func AddAgePlus(pp *Person) {
	pp.Age += 1
}
