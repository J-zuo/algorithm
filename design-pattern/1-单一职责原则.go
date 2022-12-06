package main

import "fmt"

////穿衣服的方式
//type Clothes struct {
//	//工作穿衣服
//	//逛街穿衣服
//}
//
//func (c *Clothes) OnWork() {
//	fmt.Println("工作时的装扮。。。")
//}
//
//func (c *Clothes) OnShop() {
//	fmt.Println("购物时的装扮。。。")
//}
//
//func main() {
//	c := Clothes{}
//	c.OnWork()
//	c.OnShop()
//}

// 类的职责尽量设置成单一

type ClothesShop struct{}

func (cs *ClothesShop) Style() {
	fmt.Println("逛街的装备")
}

type ClothesWork struct{}

func (cw *ClothesWork) Style() {
	fmt.Println("工作得的装备")
}

func main() {
	cw := ClothesWork{}
	cs := ClothesShop{}

	cw.Style()
	cs.Style()
}
