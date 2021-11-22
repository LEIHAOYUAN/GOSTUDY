package main

import "fmt"

func function(s1 []int) {
	s1[0] += 1000
}

func main() {

	var a = [5]int{1, 2, 3, 4, 5}

	var s = a[:]

	function(s)

	fmt.Println("切片s的第0个元素：", s[0])

	fmt.Println("原数组：", a)

}
