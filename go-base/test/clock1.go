package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func NowTime(str string) string {
	str = strings.Replace(str, "Y", "2006", -1)
	str = strings.Replace(str, "y", "06", -1)
	str = strings.Replace(str, "m", "01", -1)
	str = strings.Replace(str, "d", "02", -1)
	str = strings.Replace(str, "h", "15", -1)
	str = strings.Replace(str, "i", "04", -1)
	str = strings.Replace(str, "s", "05", -1)
	fmt.Println(str)
	return time.Now().Format(str)
}

func main() {

	str := "Y-m-d h:i:s"
	//str := "ymdhis"
	fmt.Println(NowTime(str))

	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 10000; i++ {
		fmt.Println("开始等待连接：", i+1)
		conn, err := listen.Accept()
		fmt.Println("阻塞")
		if err != nil {
			log.Println(err)
			continue
		}
		go handConn(conn)
	}
}

func handConn(conn net.Conn) {
	defer conn.Close()
	i := 0
	for {
		fmt.Println("开始处理链接连接:", i)
		_, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			return
		}
		i++
		time.Sleep(200 * time.Millisecond)
	}
}
