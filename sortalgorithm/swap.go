package sortalgorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

//异或运算交换值 （不推荐）

func SwapBit(a, b int) (c, d int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	c = a
	d = b
	return c, d
}

func SwapVar(a, b int) (c, d int) {
	var temp int
	temp = a
	c = b
	d = temp
	return c, d
}

func mergeKLists(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	n := len(lists)
	for i := 0; i < n; i++ {
		if i == 0 {
			pre = lists[i]
			continue
		}
		cur = lists[i]
		pre = merge(pre, cur)
	}
	return pre
}

func merge(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}
	}
	return head.Next
}
