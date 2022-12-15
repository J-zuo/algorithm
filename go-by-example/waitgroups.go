package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	fmt.Println("全部执行完毕。。。。")
	wg.Wait() //等待所有
	fmt.Println("程序结束")

}

func worker(id int) {
	fmt.Printf("Worder %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
