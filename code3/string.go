package main

import "fmt"

func main() {
	s := "中国人"
	fmt.Println(len(s))
	fmt.Println(s[0], s[1])
	fmt.Println(s[:5])
	fmt.Println(s[6:])
	fmt.Println(s[:])
}
