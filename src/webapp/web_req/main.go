package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/webapp/web_req/model"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求地址后的查询字符串是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息：", r.Header)
	fmt.Fprintln(w, "请求头中的Accept-Language：", r.Header["Accept-Language"])
	fmt.Fprintln(w, "请求头中的Accept-Encoding：", r.Header.Get("Accept-Encoding"))

	//获取请求体中内容的长度
	//len := r.ContentLength
	////创建byte切片
	//body := make([]byte, len)
	////将请求体中的内容读到body中
	//r.Body.Read(body)
	//fmt.Fprintln(w, "请求体中的内容：", string(body))

	////解析表单，在调用r.Form之前必须执行该操作
	//r.ParseForm()
	////获取请求参数
	//fmt.Fprintln(w, "请求参数有：", r.Form)
	//fmt.Fprintln(w, "POST请求参数有：", r.PostForm)

	//通过直接调用FormValue方法和PostFormValue方法直接获取参数值
	fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("user"))
	fmt.Fprintln(w, "form表单中的user请求参数的值是：", r.PostFormValue("username"))
}

func testJsonRes(w http.ResponseWriter, r *http.Request){
	//设置响应内容的类型
	w.Header().Set("Content-Type", "application/json")
	//创建Student
	stu := model.Students{
		ID:          1,
		Studentname: "monster",
		Password:    "123456",
		Email:       "admin@126.com",
	}
	//将stu转换为Json格式
	jsonstu, _ := json.Marshal(stu)
	//将json格式数据响应给客户端
	w.Write(jsonstu)

}

func testRedire(w http.ResponseWriter, r *http.Request){
	//设置响应头中的Location属性
	w.Header().Set("Location","https://www.baidu.com")
	//创建响应的状态码
	w.WriteHeader(302)
}


func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testJsonRes)
	http.HandleFunc("/testRedirect", testRedire)

	http.ListenAndServe(":8080", nil)
}
