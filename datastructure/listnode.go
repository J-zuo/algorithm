package main

import "fmt"

/**
判断是否为空
获取链表长度
在头部插入元素
在尾部插入元素
删除指定位置元素（第i个结点，这里注意i到底是从0开始还是从1开始计算，本题假设从1开始计位置）
删除指定值的元素（遍历链表，符合条件后删除结点，注意考虑首尾结点）
查找是否包含指定值
查找指定位置元素的值
遍历链表所有结点
*/

// 单个结点定义

type SingleListNode struct {
	data interface{}
	next *SingleListNode
}

//单链表的定义

type List struct {
	length   int
	headNode *SingleListNode
}

/*
链表的初始化
1、生成新的结点作为头结点，用头指针指向头结点
2、头结点的指针域为null
*/

func InitList() *List {
	node := new(SingleListNode)
	L := new(List)
	L.headNode = node //新建的结点就是L的头结点，此时，length为默认值0,为下面的IsNull做判断
	return L
}

/*
头插法插入结点
*/

func (list *List) AddElem(v interface{}) {
	node := &SingleListNode{data: v}
	if list.IsNull() { //处理空表的插入，否则会导致一个空的头结点后移
		list.headNode = node
		list.length++
		return
	}
	node.next = list.headNode
	list.headNode = node
	list.length++
	return
}

/**
尾插法
*/

func (list *List) AppendElem(v interface{}) {
	node := &SingleListNode{data: v}
	if list.IsNull() { //处理空表的插入，否则会导致一个空的头结点后移
		list.headNode = node
	} else {
		cur := list.headNode
		for cur.next != nil { //循环到尾部结点
			cur = cur.next
		}
		cur.next = node
	}
	list.length++
	return
}

func (list *List) IsNull() bool {
	if list.length == 0 {
		return true
	} else {
		return false
	}
}

func (list *List) ShowList() {
	if !list.IsNull() {
		cur := list.headNode
		for {
			fmt.Printf("%v", cur.data)
			fmt.Printf("-->")
			if cur.next != nil {
				cur = cur.next
			} else {
				fmt.Println("null")
				break
			}
		}
	}
	//fmt.Println(list.length)
}

// 删除第index位置的结点，index从1开始算

func (list *List) DeleteElem(index int) {
	//判断删除的位置是否合理
	if index <= 0 || index > list.length {
		fmt.Println("删除该链表位置非法")
	}
	pre := list.headNode
	//fmt.Println(pre.data)
	if index == 1 {
		list.headNode = pre.next
	} else {
		for i := 1; i < index-1; i++ {
			pre = pre.next
		}
		pre.next = pre.next.next
	}
	list.length--
}

//删除指定值的结点（一种是没有此值，一种是有一个，一种是有多个）

func (list *List) RemoveElem(v interface{}) {
	//遍历链表
	pre := list.headNode
	if pre.data == v {
		list.headNode = pre.next
	} else {
		for pre.next != nil {
			if pre.next.data == v {
				pre.next = pre.next.next
				return //用return只反回第一个就退出遍历，只找到了一个符合的元素
			} else {
				pre = pre.next
			}
		}
		fmt.Println("Fail")
		return
	}
}

func main() {
	L := InitList()
	//L.ShowList()
	L.AppendElem(1)
	L.AppendElem(2)
	L.AppendElem(3)
	L.AddElem(6)
	L.AddElem(5)
	L.AddElem(4)
	L.ShowList()
	L.DeleteElem(2)
	L.ShowList()
	L.RemoveElem(5)
	L.ShowList()
	L.RemoveElem(6)
	L.ShowList()
	L.RemoveElem(7)
	L.ShowList()

}
