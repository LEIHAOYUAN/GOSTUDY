package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := "你好，世界！"
	fmt.Fprintf(w, "%s", s)
	log.Println(s)
}

func main() {
	fmt.Println("服务启动.......")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
