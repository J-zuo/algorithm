package main

import (
	"algorithm/sortalgorithm"
	"fmt"
)

func main() {
	a, b := 1, 8
	a, b = sortalgorithm.SwapBit(a, b)
	fmt.Println(a, b)
	a, b = sortalgorithm.SwapVar(a, b)
	fmt.Println(a, b)

	//冒泡排序
	//选择排序
	//插入排序
	//归并排序
	//快速排序
	//堆排序
	//桶排序

}
