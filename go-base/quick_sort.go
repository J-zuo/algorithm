package main

import "fmt"

//快排（它是一种不稳定的排序）平均时间复杂度：nlog2n,最坏n^2
//核心思想是什么？时间复杂度和空间复杂度多少？
//步骤：
//1 找出基准条件
//2 选择一个基准元素
//3 将剩余的元素分区，左区域大于或小于基准元素，右区域小于或大于基准元素
//4 然后利用递归将左右两边执行下去，最后合并3个值
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
	var left, right []int
	//遍历比较之后分区
	for _, v := range s[1:] {
		if v >= privot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(myQuickSort(left), privot), myQuickSort(right)...) //注意 ...的用法

	//return append(myQuickSort(left), append([]int{privot}, myQuickSort(right)...)...)
}
