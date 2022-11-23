package main

import "fmt"

type Door struct {
}

func (d Door) Up() {
	fmt.Println("up")
}

func (d Door) Down() {
	fmt.Println("down")
}

//利用map和空结构体实现一个Set集合的相关操作
//
func main() {
	//空结构体的作用
	//可以实现 Set 集合
	s := make(Set)
	fmt.Println(s.Has("ddd"))
	s.Add("ddd")
	fmt.Println(s.Has("ddd"))
	s.Delete("ddd")
	fmt.Println(s.Has("ddd"))

	s2 := Door{}
	s2.Up()
	s2.Down()
}

type Set map[string]struct{}

//添加key
func (s Set) Add(key string) {
	s[key] = struct{}{}
}

//删除key
func (s Set) Delete(key string) {
	delete(s, key)
}

//查看元素在不在集合中
func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}
