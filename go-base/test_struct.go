package main

import "fmt"

type Student struct {
	Name  string
	grade int
	score int
}

func (s *Student) show() {
	fmt.Println(s)
}

func (s *Student) exam() {
	fmt.Println("学生有考试需求")
}

func main() {
	//结构体的定义与使用
	//1. 定义并初始化
	var stu Student = Student{
		Name:  "zhnagsan",
		grade: 3,
		score: 98,
	}
	fmt.Println("stu = ", stu)
	stu.exam()
	fmt.Println("-----------")
	//2.第二种方式，分开使用
	var stu1 Student
	stu1.Name = "lisi"
	stu1.grade = 4
	fmt.Println("stu1 = ", stu1)
	stu1.exam()
	fmt.Println("-----------")
	stu1.show()
	stu1.Name = "kifu"
	stu1.show()

	fmt.Println("---------")
	//第三种方式:=
	stu2 := Student{
		Name:  "hahh",
		grade: 343,
		score: 123,
	}
	fmt.Println("stu2 = ", stu2)
	fmt.Printf("stu2 type = %T,stu2 value = %v", stu2, stu2)
	stu2.show()

}
