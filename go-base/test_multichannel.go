package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 5)
	//ch2 := make(chan int, 5)

	go func() {
		defer close(ch1) //执行结束就关闭ch1
		for i := 0; i < 10; i++ {
			ch1 <- i
			//time.Sleep(1 * time.Second)
		}
	}()

	//go func() {
	//	defer close(ch2) //执行结束就关闭ch1
	//	for i := 11; i < 20; i++ {
	//		ch2 <- i
	//		//time.Sleep(1 * time.Second)
	//	}
	//}()

	fmt.Println(len(ch1), cap(ch1))

	for {
		v, ok := <-ch1
		if !ok {
			break
		}
		fmt.Println(v, ok)

		//v2, ok := <-ch2
		//if !ok {
		//	break
		//}
		//fmt.Println(v2, ok)

		//time.Sleep(1 * time.Second)
	}

}
