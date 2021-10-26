package main

import (
	"html/template"
	"net/http"
)
//IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request){
	//解析模板
	t := template.Must(template.ParseFiles("index.html"))
	//执行
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/main", IndexHandler)

	http.ListenAndServe()
}