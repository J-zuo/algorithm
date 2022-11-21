package main

import (
	"fmt"
	"sync"
)

var s []int

func worker(i int) {
	s = append(s, i+1)
	fmt.Println("正在处理 ", i, " 业务...,结果为", i+1)
}

//wait
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			worker(i)
			wg.Done()
		}(i)
	}
	wg.Wait() //等待上面的5个协程都结束再继续往后执行

	fmt.Println("程序继续执行，，")

	wg.Add(1)
	go func() {
		worker(8)
		wg.Done()
	}()
	wg.Wait()

	fmt.Println(s)
}
