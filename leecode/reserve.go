package main

import (
	"algorithm/mysdk/list"
	"fmt"
)

//翻转链表
func main() {
	L := list.InitAttachedHeadNodeList()
	L.AddHead(5)
	L.AddHead(7)
	L.AddHead(0)
	L.AddHead(12)
	L.AddHead(9)
	L.AddHead(40)
	L.PrintList()
	head := L.Head().Next //原本是返回了链表的头结点，这里使用NEXT把头去掉
	head = head.Reserve() //
	head.PrintBySlNode()  //
	fmt.Println("《-----------------》")
	L.Reserve()
	L.PrintList()

}
