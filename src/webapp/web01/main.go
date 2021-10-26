package main

import (
	"fmt"
	"net/http"
)
//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request){
	_, err := fmt.Fprintln(w, "say hello with mux!", r.URL.Path)
	if err != nil{
		fmt.Println(err)
	}
}
func main() {
	//创建一个多路复用器
	mux := http.NewServeMux()

	//将handler函数转换成处理器
	//http.HandleFunc("/", handler)
	mux.HandleFunc("/", handler)

	//创建路由，	使用默认的多路复用器
	//err := http.ListenAndServe(":8080", nil)
	//使用自己创建的多路复用器
	err := http.ListenAndServe(":8080", mux)
	if err != nil{
		fmt.Println(err)
	}
}
