go语言十大排序算法
==
## 冒泡排序
核心思想：<br>1. 外层循环次数由元素个数来确定<br>2. 内层是采用相邻元素比较还是非相邻元素<br>
代码如下：
```go
package main

import "fmt"

//假设排序的要求都是从小到大
var sl = []int{1, 20, -3, 46, 78, 27, 90, 10, 26}

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

func main() {
	fmt.Println("排序前sl: ", sl)
	//第一版排序，自己理解的冒泡排序
	//bubble1(sl)
	//bubble2(sl)
	bubble3(sl)
	//切片翻转
	reverse(sl)

}
```

## 快排
介绍：快排是比较优秀的排序算法,
## 插入排序
介绍：一般我们指的是*简单插入排序*或*直接插入排序*,时间复杂度是n^2，它是一种稳定排序
核心思想：每次把一个数插到已经排好序的数列里面形成新的排好序数列<br>
理解方式：想象成打牌
## 选择排序
介绍：选择排序也类似于打牌，通过内层循环完毕寻找到最小值下标，然后进行替换， <br>
每一次大循环都能将最小值放在开头位置<br>
注意：程序中注意需要拿第一个作为默认最小值，然后其他下标对应值与默认最小值进行比较
## 归并排序
介绍：1945年，由现代计算机之父冯诺依曼发明，它是一种分治策略的排序算法，通过递归使每个子序列有序<br>
再将两个有序的子序列合并成一个有序序列
## 优先队列堆排序
