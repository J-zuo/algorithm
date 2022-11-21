package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan int, 10)
	var wg sync.WaitGroup
	//开启N个后台打印go程
	for i := 0; i < cap(done); i++ {
		//fmt.Println("for 循环次数：", i)
		wg.Add(1)
		go func() {
			fmt.Println("写入：")
			done <- 1
			wg.Done()
		}()
	}
	wg.Wait()
	for i := 0; i < cap(done); i++ {
		fmt.Println(<-done)
	}

}
