package list

import "fmt"

// 定义单向链表中的一个结点
type SLNode struct {
	Data int
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
func (l *SList) PrintList() {
	if l.lenth > 0 { //说明链表不为空
		cur := l.head.Next
		for {
			fmt.Printf("%v", cur.Data)
			fmt.Printf("->")
			if cur.Next != nil {
				cur = cur.Next
			} else {
				fmt.Println("null")
				break
			}
		}

	} else {
		fmt.Println(l.lenth, l.head.Data, l.head.Next) //初始化链表之后，不做任何操作，打印该链表： 0 0 <nil>
	}

}

func (l *SList) Head() *SLNode {
	return l.head
}

func (l *SList) Reserve() *SLNode {
	var pre *SLNode
	cur := l.head.Next

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

//根据首结点打印出链表(方便展示)

func (n *SLNode) PrintBySlNode() {
	cur := n
	for {
		fmt.Printf("%v", cur.Data)
		fmt.Printf("->")
		if cur.Next != nil {
			cur = cur.Next
		} else {
			fmt.Println("null")
			break
		}
	}
}

func (n *SLNode) Reserve() *SLNode {
	var pre *SLNode
	cur := n
	fmt.Println(cur.Data)
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
