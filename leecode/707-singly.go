package main

import "fmt"

type ListNode struct {
	data int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	size int
}

// 初始化链表

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{}, 0}
}

//根据下标来插入结点(下标从0开始计)

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = max(index, 0)
	fmt.Printf("%v\n", index)
	l.size++
	pred := l.head
	fmt.Println("头结点的默认值：", pred.data)
	for i := 0; i < index; i++ { //找到原来下标为index的结点的前驱结点
		pred = pred.Next
	}
	toAdd := &ListNode{val, pred.Next} //创建新结点toAd，将toAdd的后继结点设为pred的后继结点，
	pred.Next = toAdd                  //将pred的后继结点更新为toAdd，这样就完成了插入

}

//头插法

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

//尾插法

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

//链表打印

func (l *MyLinkedList) PrintList() {
	if l.size > 0 {
		cur := l.head
		for {
			fmt.Printf("%v", cur.data)
			fmt.Printf("-->")
			if cur.Next != nil {
				cur = cur.Next
			} else {
				fmt.Println("null")
				break
			}
		}
	} else {
		fmt.Println("linked list is null")
	}
}

//获取指定下标元素

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index > l.size {
		return -1
	}
	cur := l.head
	for i := 0; i <= index; i++ { //这里是小于等于，如果index=0,就是第一个元素
		cur = cur.Next
	}
	fmt.Printf("Get 获得到的值为:%v\n", cur.data)
	return cur.data
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func (l *MyLinkedList) DeleteAtIndex(index int) { //删除的是下标
	if index < 0 || index >= l.size { //这里的等于号是长度大于下标，是非法的
		return
	}
	l.size--
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}

func main() {
	L := Constructor()
	L.PrintList()
	L.AddAtHead(5)
	L.AddAtHead(9)
	L.AddAtIndex(2, 3)
	L.PrintList()
	L.AddAtTail(12)
	L.PrintList()
	L.DeleteAtIndex(3)
	L.PrintList()
	L.Get(0)

}
