package main

import "fmt"

func main() {
	s := "leihaoyuan"
	fmt.Println("字符串长度：", len(s))
	fmt.Println(s[0], s[1])
	fmt.Println("字符串截取，到index = 4的位置结束：", s[:5])
	fmt.Println("字符串截取，从index = 6的位置开始：", s[6:])
	fmt.Println("截取全部字符串：", s[:])
}
