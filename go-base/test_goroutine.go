package main

import (
	"fmt"
	"log"
	"net/http"
)

//请对你创建的goroutine负责

func setup() {
	fmt.Println("程序初始化操作")
}

func main() {
	//初始化
	setup()
	//主服务
	go server()

	//for debug
	go pprof()

	select {}
}

func pprof() {
	//另起一个goroutine来监听8081端口，用于debug
	mux := http.NewServeMux()
	mux.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("debug")) //向客户端发送该字符串
	})
	if err := http.ListenAndServe(":8082", mux); err != nil {
		fmt.Println(err)
	}
}

func server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong")) //向客户端发送该字符串
	})

	//主服务
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panicln("http server err:", err)
		return
	}
}
