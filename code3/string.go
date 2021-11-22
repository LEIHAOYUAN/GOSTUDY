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

	// 其他字符串API
	// strings:提供搜索、比较、切分与字符串连接功能
	// bytes:对字符串的底层字节进行操作，可以使用[]bytes转换类型后进行处理
	// strconv:主要是字符串与其他类型的转换，比如整数、布尔类型
	// unicode:主要是对字符串中的单个字符做判断，比如IsLetter、IsDigit、IsUpper

}
