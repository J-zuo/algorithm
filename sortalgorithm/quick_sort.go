package sortalgorithm

import (
	"fmt"
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
	fmt.Println(s1 == nil)
	for i := 0; i < 1; i++ {
		s1 := append(s1, rand.Intn(1000))
		fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", s1, &s1, len(s1), cap(s1))
		fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", s1, &s1, len(s1), cap(s1))
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
