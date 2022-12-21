package main

import "fmt"

//假设排序的要求都是从小到大
var sl = []int{1, 20, -3, 46, 78, 27, 90, 10, 26}

//var sl = []int{5, 4, 3, 2, 1}
//var sl = []int{1, 2, 3, 4, 5}

//var sl = []int{-3, 1, 10, 20, 26, 27, 46, 78, 90}

//bubble1 冒泡排序1,也是最好理解的版本
func bubble1(data []int) {
	//获取切片的长度，用于外层循环
	lenth := len(sl)

	//开始外层循环，外层每一次循环结束，将最大的放在最后面
	for i := 0; i < lenth; i++ {
		for j := 0; j < lenth-i-1; j++ {
			//如果左边的元素大于右边，就交换
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	fmt.Println("排序后sl: ", data)
}

func bubble2(data []int) {
	lenth := len(data)
	for i := 0; i < lenth; i++ { //这种方式循环的话，每一次外层循环结束，把当前层的最小值放在前面
		for j := i + 1; j < lenth; j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i] //不是相邻元素比较，而是和每一外层固定的元素进行比较，
			}
		}
	}
	fmt.Println("排序后sl: ", data)
}

func bubble3(data []int) {
	lenth := len(data)
	for i := 0; i < lenth; i++ { //每一层都把最小的放在前面，和bubble1正好是反过来的
		flag := false                     //这种优化避免无效的比较，当内层循环没有形成交换意味着有序，其他外层无序比较了
		for j := lenth - 2; j >= i; j-- { //从后往前
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j] //是相邻两个元素进行比较
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("排序后sl: ", data)
}

func reverse(data []int) {
	lenth := len(data)
	for i := 0; i < lenth/2; i++ {
		data[i], data[lenth-i-1] = data[lenth-i-1], data[i]
	}
	fmt.Println("翻转之后: ", data)
}

/**
简单的插入排序，时间复杂度n^2，稳定排序
模拟打牌的方式
*/
func insertsort(data []int) {
	lenth := len(data)
	for i := 1; i < lenth; i++ { //从下标为1开始是因为我们默认第一个数是已经排好序的
		deal := data[i]     //待排序的数
		j := i - 1          //待排序的数左边的第一个数的位置
		if deal < data[j] { //如果待排序得到数小于左边的数，就需要处理
			//一直往左边找，比带排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < data[j]; j-- {
				data[j+1] = data[j] //某个数后移，留空位给待排序
			}
			data[j+1] = deal //待排序的数插入空位
			fmt.Println(data)
		}
	}
	fmt.Println("插入排序之后：", data)
}

/**
selectsort，选择排序，时间复杂度n^2，是一种不稳定的排序
也是模拟打牌的方式：打扑克牌的时候，会习惯性地从左到右扫描，然后将最小的牌放在最左边，
然后从第二张牌开始继续从左到右扫描第二小的牌，放在最小的牌右边，以此反复。
核心：1.找到下标，每个轮次只交换一次，不同于冒泡需要交换多次
写代码时要注意：1.每个轮次我们都要确定最小元素以及最小元素对应的下标，用于和内层迭代时的元素相比较
*/
func selectsort(data []int) {
	lenth := len(data)
	//进行lenth-1次迭代就行了，因为最后一个数必然是最大/最小的
	for i := 0; i < lenth-1; i++ {
		//假设每一轮次第一个元素是最小的
		min := data[i]
		//最小元素的下标
		minIndex := i
		for j := i + 1; j < lenth; j++ {
			if data[j] < min {
				min = data[j]
				minIndex = j
			}
		}
		//找到最小下标如果不等于最开始的下标，意味着里面肯定存在比设定的最小值要小，需要交换
		if i != minIndex {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
	fmt.Println("选择排序之后:", data)
}

func main() {
	fmt.Println("排序前sl: ", sl)
	//第一版排序，自己理解的冒泡排序
	//bubble1(sl)
	//bubble2(sl)
	//bubble3(sl)
	//reverse(sl)
	//insertsort(sl)
	selectsort(sl)

}
