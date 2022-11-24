package sortalgorithm

import (
	"fmt"
	"math/rand"
	"time"
)

var normal = []int{100, 50, 10, 90, 30, 70, 40, 80, 60, 20} //自定义一个切片

//定义一个含有切片的结构体
type SqList struct {
	Slice []int
	Lenth int
}

//初始化随机切片
func InitSlice(lenth int) *SqList {
	rand.Seed(time.Now().UnixNano())
	s1 := []int{}
	//fmt.Println(s1 == nil)
	for i := 0; i < lenth; i++ {
		s1 = append(s1, rand.Intn(1000))
		//fmt.Printf("Pointer = %p, len = %d, cap = %d\n", &s1, len(s1), cap(s1))
		//fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", s1, &s1, len(s1), cap(s1))
		//fmt.Println(len(s), cap(s))
	}

	return &SqList{Slice: normal, Lenth: len(normal)}
	//return &SqList{Slice: s1, Lenth: lenth} //返回程序随机生成的
}

//定义一个方法，交换切片中下标为i,j的值
func (s *SqList) Swap(i, j int) {
	temp := s.Slice[i]
	s.Slice[i] = s.Slice[j]
	s.Slice[j] = temp
}

//冒泡排序(乞丐版)(无脑版本)(假冒伪劣产品)(蹭流量)
func (s *SqList) Bubble() {
	var i, j int
	for i = 0; i < s.Lenth; i++ {
		for j = i + 1; j < s.Lenth; j++ {
			if s.Slice[i] > s.Slice[j] {
				s.Swap(i, j)
			}
		}
	}
}

//冒泡排序(正版)(普通版)(国家队)(是真冒泡)
func (s *SqList) Bubble2() {
	var i, j int
	for i = 0; i < s.Lenth; i++ {
		for j = s.Lenth - 2; j >= i; j-- {
			if s.Slice[j] > s.Slice[j+1] {
				s.Swap(j, j+1)
			}
		}
	}
}

//冒泡排序(正装升级版)
func (s *SqList) Bubble3() {
	var i, j int
	for i = 0; i < s.Lenth; i++ {
		flag := false
		for j = s.Lenth - 2; j >= i; j-- {
			if s.Slice[j] > s.Slice[j+1] {
				s.Swap(j, j+1)
				flag = true //标记交换就置为true,
			}
		}
		if !flag {
			break
		}
	}
}

//快速排序(普通版)
func (s *SqList) QuickSort() {
	Qsort(s, 0, s.Lenth-1)
}

func Qsort(s *SqList, low int, high int) {
	//确立分区
	if low < high {
		pivot := Partition(s, low, high)
		Qsort(s, low, pivot-1)
		Qsort(s, pivot+1, high)
	}
}

//选定一个位置，使其对应的值比左边都大，比右边都小
func Partition(s *SqList, low, high int) int {
	fmt.Println(s, low, high)
	pivotkey := s.Slice[low] //用第一个值作为枢轴 ,在交换过程中不断变换自己的位置，接下来就是递归
	fmt.Println(pivotkey)
	for low < high {
		for low < high && s.Slice[high] >= pivotkey {
			high--
		}
		s.Swap(low, high)
		for low < high && s.Slice[low] <= pivotkey {
			low++
		}
		s.Swap(low, high)
	}
	fmt.Println(low)
	return low
}
