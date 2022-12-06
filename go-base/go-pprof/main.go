package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len:%d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10) //10毫秒
		}
	}()
	http.ListenAndServe("127.0.0.1:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	fmt.Println(string(data))
	datas = append(datas, string(data))
	return len(datas)
}
