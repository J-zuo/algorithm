package main

import "fmt"

// 这个概念很重要，开闭原则。
// 类的改动是通过增加代码来改动的，而不是修改之前的代码

type Banker struct {
}

func (b *Banker) Save() {
	fmt.Println("进行了存款业务。。。。")
}

func (b *Banker) Transfer() {
	fmt.Println("进行了转账业务。。。。")
}

func (b *Banker) Pay() {
	fmt.Println("进行了付款业务。。。。")
}

//需要另加
func (b *Banker) Shares() {
	fmt.Println("进行了股票业务。。。。")
}

func main() {
	b := Banker{}
	b.Save()
	b.Transfer()
	b.Pay()
	//如果再增加一个功能，是不是就需要再次定义一个方法,这就违背了开闭原则
	b.Shares()
}
