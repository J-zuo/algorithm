package main

import (
	"fmt"
	"strconv"
)

func main() {

	//----------数字字符串转为各种int类型----------
	string := "12322"
	//string到int
	a, _ := strconv.Atoi(string)
	//string到int64
	b, _ := strconv.ParseInt(string, 10, 64)
	c, _ := strconv.ParseUint(string, 10, 64)
	fmt.Printf("%T-%T-%T,%v-%v-%v\n", a, b, c, a, b, c)

	//----------各种int类型转为数字字符串----------
	//int到string
	d := 124
	e := strconv.Itoa(d)
	//int64到string
	f := strconv.FormatInt(int64(d), 10)
	fmt.Printf("%T-%T,%v-%v\n", e, f, e, f)

	//string到float32(float64)
	g := "23.4"
	float, _ := strconv.ParseFloat(g, 64)
	fmt.Printf("%T, %v\n", float, float)
	////float到string
	//string := strconv.FormatFloat(float32, 'E', -1, 32)
	h := strconv.FormatFloat(0.1, 'E', -1, 64)
	fmt.Printf("%T, %v\n", h, h)
}
