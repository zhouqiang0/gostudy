package model

import (
	"fmt"
	"web/src/webapp/web_db/utils"
)

//students结构体 （同mysql数据库）
type Students struct {
	ID int
	Studentname string
	Password string
	Email string
}

func (user *Students) AddUser() error{
	//1.写sql语句
	sqlStr := "insert into students(studentname,password,email) value(?,?,?)"
	//2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("预编译异常 :", err)
		return err
	}
	//3.执行
	_, err = inStmt.Exec("admin","123456","admin@126.com")
	if err != nil{
		fmt.Println("执行异常 :", err)
		return err
	}
	return nil
}

//不带预编译
func (user *Students) AddUser2() error{
	//1.写sql语句
	sqlStr := "insert into students(studentname,password,email) value(?,?,?)"
	////2.预编译
	//inStmt, err := utils.Db.Prepare(sqlStr)
	//if err != nil{
	//	fmt.Println("预编译异常 :", err)
	//	return err
	//}
	//3.执行
	_, err := utils.Db.Exec(sqlStr,"admin2","666666","admin2@qq.com")
	if err != nil{
		fmt.Println("执行异常 :", err)
		return err
	}
	return nil
}
//根据用户id查询一条记录
func (stu *Students) GetById(stuid int)(*Students, error){
	//写sql
	sqlStr := "select id, studentname, password, email from students where id = ?"
	row := utils.Db.QueryRow(sqlStr, stuid)
	//声明变量
	var id int
	var studentname, password, email string
	err := row.Scan(&id, &studentname, &password, &email)
	if err != nil{
		fmt.Println("查询失败：", err)
		return nil, err
	}
	s := &Students{
		ID:          id,
		Studentname: studentname,
		Password:    password,
		Email:       email,
	}
	return s, nil
}

//获取数据库中所有记录
func (stu *Students) GetStudents()([]*Students, error){
	sqlStr := "select id, studentname, password, email from students"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil{
		fmt.Println("查询失败")
		return nil, err
	}
	var id int
	var studentname, password, email string
	//next反复查询是否有下一条记录
	//创建一个students切片记录
	var students []*Students
	for rows.Next(){
		err = rows.Scan(&id, &studentname, &password, &email)
		if err != nil{
			fmt.Println("查询失败：", err)
			return nil, err
		}
		s := &Students{
			ID:          id,
			Studentname: studentname,
			Password:    password,
			Email:       email,
		}
		students = append(students, s)
	}
	return students, nil
}
