package main

import "fmt"

func main() {
	s1 := make(map[string]string, 5)
	s1["zs"] = "zhangsan"
	s1["ls"] = "lisi"
	s1["ww"] = "wangwu"
	s1["db"] = "dabai"
	s1["hp"] = "heipi"

	fmt.Println(s1)
	s1["ww"] = "wwwwwwwww"

	fmt.Println(s1)

	a1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a1)
	a2 := a1 //发生值拷贝
	fmt.Println(a2)
	a2[2] = 8
	fmt.Println(a2, a1)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1
	fmt.Println(slice2)
	slice2[0] = 9
	fmt.Println(slice2, slice1)
}
