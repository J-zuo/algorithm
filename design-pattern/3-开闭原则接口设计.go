package main

import "fmt"

type AbstrctBanker interface {
	DoBusi() //抽象的处理业务接口
}

type SaveBanker struct {
	//AbstrctBanker
}

func (sb *SaveBanker) DoBusi() {
	//TODO implement me
	fmt.Println("进行了存款")
}

// +++++++++++++++++++++++
//转账业务
type TransferBanker struct {
	//AbstrctBanker
}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

func (tb *TransferBanker) Dobusi() {

}

//实现一个架构层(基于抽象层进行业务封装-针对interface接口进行封装)
// 多态的特点：通过多肽，你传的是子类对象，父类指针指向子类对象，通过父类指针来调用接口就能调到具体对象的具体方法
func BankBusiness(banker AbstrctBanker) {
	//通过接口向下调用（产生多态现象）
	banker.DoBusi()
}

func main() {
	//存款的业务
	//sb := SaveBanker{}
	//sb.Dobusi()
	//
	//tb := TransferBanker{}
	//tb.Dobusi()

	//类的改动是通过增加代码进行的，而不是修改源代码
	//大部分设计模式都是为了满足开闭原则，能够保证不破坏系统的情况下给系统增加功能
	BankBusiness(&SaveBanker{})
	BankBusiness(&TransferBanker{})

}
