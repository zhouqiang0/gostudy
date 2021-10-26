package model

import (
	"fmt"
	"testing"
)

//TestMain函数可以在测试函数执行之前做其他操作
func TestMain(m *testing.M)  {
	fmt.Println("测试开始：")
	m.Run()
}
func TestStudents(t *testing.T) {
	fmt.Println("子测试程序")
	//t.Run("测试添加用户", testStudents_AddUser)
	//t.Run("测试查询一条记录", testStudents_GetById)
	t.Run("测试查询所有记录：", testStudents_GetStudents)
}

func testStudents_AddUser(t *testing.T) {
	fmt.Println("测试添加用户")
	user := &Students{}
	//调用方法
	user.AddUser()
	user.AddUser2()
}

func testStudents_GetById(t *testing.T) {
	fmt.Println("测试查询一条记录：")
	stu := Students{ID:1}

	u, _ := stu.GetById(stu.ID)
	fmt.Println("查询结果：", u)
}

func testStudents_GetStudents(t *testing.T) {
	fmt.Println("查询所有记录")
	stu := &Students{}
	ss, _ := stu.GetStudents()
	//遍历切片ss
	for i, student := range ss{
		fmt.Printf("第%d个用户信息是%v：",i+1 , student)
	}
}
