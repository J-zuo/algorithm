package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//树的创建

}

func CreateTreeNode(data int) *TreeNode {
	return &TreeNode{data, nil, nil}
}
