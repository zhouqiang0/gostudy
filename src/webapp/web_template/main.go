package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request){
	//解析模板, 使用must函数自动处理错误
	tp := template.Must(template.ParseFiles("src/webapp/web_template/index1.html", "src/webapp/web_template/index2.html"))

	////执行，在index1中显示响应数据
	//err = tp.Execute(w, "123")
	//if err != nil{
	//	fmt.Println("err : ", err)
	//}
	//执行，在index2中显示响应数据
	tp.ExecuteTemplate(w, "index2.html", "转到index2.html")
}

func main() {
	http.HandleFunc("/testTemplate", testTemplate)

	http.ListenAndServe(":8080", nil)
}
