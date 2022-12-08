package main

import (
	"algorithm/sortalgorithm"
	"fmt"
)

func main() {
	//切片的拷贝
	//var s1 []int
	//s1 = []int{9, 8, 7, 10, 14, 19, 12, 12}
	//s2 := []int{1, 2, 3, 5, 6}
	//fmt.Printf("%p, %v, %d, %d\n", &s1, s1, cap(s1), len(s1))
	//fmt.Printf("%p, %v, %d, %d\n", &s2, s2, cap(s2), len(s2))
	//copy(s1, s2)
	//fmt.Printf("%p, %v, %d, %d\n", &s1, s1, cap(s1), len(s1))

	//数字转换为字符串
	//a := fmt.Sprintf("%d", 12345)
	//fmt.Printf("type a: %T\n", a)

	//生成一个随机切片
	sq := sortalgorithm.InitSlice(10)
	fmt.Println(sq.Slice, sq.Lenth)

	//--------------------冒泡----------------------------
	//sq.Swap(1, 2)
	//fmt.Println(sq.Slice, sq.Lenth)
	//sq.Bubble()
	//sq.Bubble2()
	//fmt.Println(sq.Slice, sq.Lenth)
	//fmt.Println(sq, sq.Slice, sq.Lenth)

	//--------------------快排----------------------------
	//快排的时间复杂度
	sq.QuickSort()

}
