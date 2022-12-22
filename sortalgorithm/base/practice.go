package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 5, 0, -1}
	//insertsort1(sl)
	selectsort1(sl)

}

func insertsort1(data []int) {
	//Get dynamic array lenth
	lenth := len(data)
	//外层循环，总共需要n-1次循环
	for i := 1; i < lenth; i++ {
		//待插入的数
		deal := data[i]
		//带插入数左边的第一个数的位置
		j := i - 1
		//如果待插入的数比左边的数小，就需要处理
		if deal < data[j] {
			for ; j >= 0 && deal < data[j]; j-- {
				data[j+1] = data[j]
			}
			data[j+1] = deal
		}
	}
	fmt.Println(data)
}

func selectsort1(data []int) {
	lenth := len(data)
	for i := 0; i < lenth-1; i++ {
		//每次外层循环假设第一个元素为最小值
		min := data[i]
		minIndex := i
		for j := i + 1; j < lenth; j++ {
			if data[j] < min {
				min = data[j]
				minIndex = j
			}
		}
		if i != minIndex { //如果内层处理完之后，i没有比变，说明没有比i值更小，变了就需要交换最小值
			data[i], data[minIndex] = data[minIndex], data[i]
		}

	}
	fmt.Println(data)
}
