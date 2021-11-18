package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(_ http.ResponseWriter, _ *http.Request) {
	s := "你好，世界！"
	log.Println(s)
}

func main() {
	fmt.Println("服务启动.......")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
