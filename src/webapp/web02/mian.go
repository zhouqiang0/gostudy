package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	_, err := fmt.Fprintln(w, "自己实现的handler接口,实现server结构")
	if err != nil{
		fmt.Println("err: ", err)
	}
}

func main() {
	myhandler := MyHandler{}

	//http.Handle("/myHandler", &myhandler)
	//创建Server结构
	server := http.Server{
		Addr:              ":8080",
		Handler:           &myhandler,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 2 * time.Second,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}
	//err := http.ListenAndServe(":8080", nil)
	err := server.ListenAndServe()
	if err != nil{
		fmt.Println("err: ", err)
	}
	
}
