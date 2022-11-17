package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Name   string
	Author string
	//Ibsn   *int  //point类型，在同一个struct中的两个不同的实例中，地址肯定是不一样的。所以虽然指针类型可以比较，但是比较结果是false
	Asc []int //因为此类型是slice,此结构体的两个实例是不能比较的，会检查不通过[invalid operation: b1 == b2 (struct containing []int cannot be compared)]
}

type Ebook struct {
	Name   string
	Author string

	Asc []int
}

//结构体之间比较的问题
//go语言中不能比较的类型有slice map func
func main() {
	//第一种情况，同一struct的两个实例能不能比较
	b1 := Book{"三国志", "承受", []int{1, 2, 3}}
	fmt.Println(b1)

	b3 := Book{"三国志", "承受", []int{1, 2, 3}}
	fmt.Println(b3)

	b2 := Book{"水浒传", "斯乃安", []int{1, 2, 3}}
	fmt.Println(b2)

	//if b1 == b2 { //如果struct 属性中有不可比较的类型，直接比较是编译通不过的
	if reflect.DeepEqual(b1, b3) { //使用reflect反射中的方法进行深度比较,是一种比较的方式
		fmt.Printf("b1 type:%T, value:%v\n", b1, b1)
	} else {
		fmt.Println("不相等")
	}

	//fmt.Println("--------")
	//if b1 == b3 {
	//	fmt.Printf("b1 type:%T, value:%v", b1, b2)
	//} else {
	//	fmt.Println("不相等")
	//}

	fmt.Println("-----------")
	//第二种情况，两个不同struct实例之间的比较

	//b4 := Ebook{
	//	Name:   "ebook1",
	//	Author: "dfk",
	//}
	//b5 := Book{
	//	Name:   "ebook1",
	//	Author: "dfk",
	//}

	//if b4 == b5 { //不同的struct之间不能比较，即使内部的属性相同也不能直接==比较
	//
	//}

	//强制类型转换进行比较
	var b6 Human
	var b7 Superman

	b8 := Human(b7)               //我把b7强制转换成Human类型。
	fmt.Println(b6 == b8, b6, b7) //这种就是可以比较的，前提是struct属性是可比较的且属性相同

	fmt.Println("---------")

	b9 := Human1{"zuo", 12, new(int)} //如果是已经初始化了的，这样可以直接比较，但是不相同
	b10 := Superman1{"zuo", 12, new(int)}
	b11 := Human1(b10)
	fmt.Println(b11 == b9, b11, b9)

}

type Human struct {
	Name string
	age  int
	sex  *int
}

type Superman struct {
	Name string
	age  int
	sex  *int
}

type Human1 struct {
	Name string
	age  int
	sex  *int
}

type Superman1 struct {
	Name string
	age  int
	sex  *int
}
