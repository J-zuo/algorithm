package main

import "fmt"

//司机张三，李四， 汽车宝马，奔驰

//1 -- 张三开奔驰
//2 -- 李四开宝马

//奔驰汽车
type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running...")
}

//宝马汽车
type Bmw struct {
}

func (b *Bmw) Run() {
	fmt.Println("Bmw is running...")
}
