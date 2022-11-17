package main

import "fmt"

/**
Leecode 206:Reverse Linked List

Given the head of singly linked list, reverse the list, and return reversed list.

*/

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

type MyLinkedList1 struct {
	size int
	head *ListNode1
}

func Constructor1() MyLinkedList1 {
	return MyLinkedList1{0, &ListNode1{}}
}

func (l *MyLinkedList1) addIndex(index int, val int) {
	if index > l.size {
		return
	}
	index = max1(index, 0)
	l.size++
	pred := l.head
	for i := 0; i <= index; i++ {
		pred = pred.Next
	}
	toAdd := &ListNode1{val, pred.Next}
	pred.Next = toAdd
}

func (l *MyLinkedList1) addHead(val int) {
	l.addIndex(0, 3)
}

func (l *MyLinkedList1) PrintList() {
	if l.size > 0 {
		cur := l.head
		for {
			fmt.Printf("%v", cur.Val)
			fmt.Printf("-->")
			if cur.Next != nil {
				cur = cur.Next
			} else {
				fmt.Println("null")
				break
			}
		}
	} else {
		fmt.Println("Linked list size 0")
	}
}

func max1(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {

	//1、初始化带头结点的链表，有效长度为0
	//2、利用头插法或者是尾插法来生产一个有长度的链表
	//3、实现翻转链表的函数，并调用，看结果
	L := Constructor1()
	L.addIndex(0, 3)
	L.PrintList()

}
