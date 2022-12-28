package main

import "fmt"

func main() {
	//利用结构体来实现集合，go语言本身没有自带的集合结构，可以自己实现
	//集合的的特征：不重复，利用map和空结构体来实现
	//集合的方法：1，添加元素 2，判断元素是否存在 3，删除元素

	v1 := make(MySet)
	v1.AddElem("zuo")
	v1.AddElem("zuo1")
	v1.AddElem("zuo2")
	v1.AddElem("zuo3")
	v1.AddElem("zuo4")
	fmt.Println(v1)
	fmt.Println(v1.ElemIsTrue("zuo8"))
	fmt.Println(v1.ElemIsTrue("zuo3"))
	v1.DelElem("zuo")
	fmt.Println(v1)

}

type MySet map[string]struct{}

func (ms MySet) AddElem(key string) {
	ms[key] = struct{}{}
}

func (ms MySet) ElemIsTrue(key string) bool {
	_, ok := ms[key]
	if ok {
		return true
	}
	return false
}

func (ms MySet) DelElem(key string) {
	delete(ms, key)
}
