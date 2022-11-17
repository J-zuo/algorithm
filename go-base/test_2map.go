package main

import (
	"fmt"
	"sort"
)

func PrintIntStringMap(m map[int]string) {
	for k, v := range m {
		fmt.Printf("key = %v, val = %v\n", k, v)
	}

}

//如何将map中的key进行排序
func main() {
	m1 := make(map[int]string)
	m1[0] = "a"
	m1[1] = "b"
	m1[2] = "c"
	m1[3] = "d"
	m1[4] = "e"
	m1[5] = "f"
	m1[6] = "g"
	m1[7] = "h"
	m1[8] = "i"
	m1[9] = "j"
	m1[10] = "k"
	m1[11] = "l"
	m1[12] = "m"
	m1[13] = "n"

	PrintIntStringMap(m1)
	//上述结果可知，输出的map是无序的
	//变成有序的
	fmt.Println("------")
	fmt.Println(SortMapByKey(m1))

}

//根据map中的key来排序输出map

func SortMapByKey(m map[int]string) map[int]string {
	//第一步：使用切片来保存map的key
	s := []int{}
	for k, _ := range m {
		s = append(s, k)
	}
	//第二步：将[]int类型的切片进行升序排序
	sort.Ints(s)
	//声明一个新的map，用来存储新排序好的目标map
	m2 := make(map[int]string)
	//第三步：遍历排序好的切片,然后按切片的值对应查找原map中的值并赋值新的map
	for _, v := range s {
		m2[v] = m[v]
	}
	//fmt.Println(m2)
	return m2
}
