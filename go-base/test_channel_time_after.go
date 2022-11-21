package main

import (
	"fmt"
	"time"
)

func main() {

	//ch := make(chan int)

	//go func() {
	//	time.Sleep(3 * time.Second)
	//	ch <- 1
	//}()

	for {
		select {
		//case res := <-ch:
		//	fmt.Println(res)
		case <-time.After(time.Second * 10):
			fmt.Println("时间到了")
			break
		case <-time.Tick(time.Second * 1):
			fmt.Println("tick!")
		}
	}

	fmt.Println("结束主gorotine")

}
