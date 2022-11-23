package sortalgorithm

import (
	"math/rand"
	"time"
)

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
	return &SqList{Slice: s1, Lenth: lenth}
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
