package main

import "fmt"

func main() {

	//定义一个变量，但是没有分配内存空间,此时p,q都为nil
	var p *int
	var q *int
	fmt.Println(p == q, p, q)

	//给p,q分配内存空间，他的值是地址，
	p = new(int)
	q = new(int)

	*p = 10

	fmt.Println(p == q, p, q, *p, *q, &p, &q)

}
