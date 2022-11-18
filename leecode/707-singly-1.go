package main

import "fmt"

// 定义单向链表中的一个结点
type SLNode struct {
	data int
	Next *SLNode
}

// 定义单向链表的
type SList struct {
	head  *SLNode //取名很重要，这里取head帮助理解，表示这是一个头结点，数据类型是
	lenth int
}

//初始化一个带头结点的链表
func InitAttachedHeadNodeList() *SList {
	return &SList{&SLNode{}, 0} //链表长度为0,因为头结点不算
}

func main() {
	fmt.Println("单向链表的基本操作")

	//实例化一个链表
	L := InitAttachedHeadNodeList()

	//模拟打印一下初始化的链表，方面理解
	//L.printList()

	//根据元素的下标（从0开始计）来插入结点，分三种情况：
	//1、头插法其实就是找到下标为0的结点的前驱结点，正好就是我们初始化的头结点。
	//2、尾插法就是找到长度为下标为lenth的前驱几点（也就是未插入前最后一个元素，下标为lenth-1），加上待插入的结点之后，下标正好是lenth
	//3、正常的中间插入就是找到下标为n(0<=n<lenth)
	L.AddHead(12)
	//L.printList()
	L.AddHead(8)
	L.printList()

}

func (l *SList) AddIndex(index int, val int) {
	if index > l.lenth {
		return
	}
	if index < 0 {
		index = 0
	}
	//找到下标对应结点的前驱结点
	pred := l.head //定义一个初始指针
	//遍历查找前驱结点(这里需要注意空表时第一个插入结点)
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	//程序到此时就找到了插入准确位置的条件
	//申请一个结点
	newNode := &SLNode{val, pred.Next}
	pred.Next = newNode
	l.lenth++
}

func (l *SList) AddHead(val int) {
	l.AddIndex(0, val)
}

func (l *SList) AddTail(val int) {
	l.AddIndex(l.lenth, val)
}

//传入原链表的头指针，打印链表
func (l *SList) printList() {
	if l.lenth > 0 { //说明链表不为空
		cur := l.head.Next
		for {
			fmt.Printf("%v", cur.data)
			fmt.Printf("->")
			if cur.Next != nil {
				cur = cur.Next
			} else {
				fmt.Println("null")
				break
			}
		}

	} else {
		fmt.Println(l.lenth, l.head.data, l.head.Next) //初始化链表之后，不做任何操作，打印该链表： 0 0 <nil>
	}

}
