package main

import "fmt"

func main() {
	s := "leihaoyuan"

	fmt.Println("字符串s字节数：", len(s))
	fmt.Println("下标index：0=", s[0], "下标index：1=", s[1])
	fmt.Println("字符串截取，到index = 4的位置结束：", s[:5])
	fmt.Println("字符串截取，从index = 6的位置开始：", s[6:])
	fmt.Println("截取全部字符串：", s[:])

	fmt.Println(".............................")

	var s1 = s
	s1 = "abcdefghijklmn"
	fmt.Println("修改后的s1=", s1)

	fmt.Println(".............................")

	fmt.Println("字符串s1字节数：", len(s1), "s1的地址：", &s1)

	fmt.Println("字符串s字节数：", len(s), "s的地址：", &s)
	fmt.Println("s[0]=", s[0], "s[1]=", s[1])
	var sub1 = s[:5]
	fmt.Println("字符串截取，到index = 4的位置结束：", sub1, "s[:5]地址：", &sub1)
	var sub2 = s[6:]
	fmt.Println("字符串截取，从index = 6的位置开始：", sub2, "s[6:]地址：", &sub2)
	var sub3 = s[:]
	fmt.Println("截取全部字符串：", sub3, "s[:]地址：", &sub3)

}
