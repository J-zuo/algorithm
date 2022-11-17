package sortalgorithm

import "fmt"

//快排
//核心思想是什么？时间复杂度和空间复杂度多少？
//步骤：
func main() {
	//定义一[]int切片
	s := []int{32, 4, -9, 20, 99, 40, 3}
	a := myQuickSort(s)
	fmt.Println(a)
}

func myQuickSort(s []int) []int {
	//递归结束条件
	if len(s) < 2 {
		return s
	}
	//选择一个基准元素值
	privot := s[0]
	left := []int{}
	right := []int{}
	//遍历比较之后分区
	for _, v := range s {
		if v <= privot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(myQuickSort(left), privot), myQuickSort(right)...)
}
