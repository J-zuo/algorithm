package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const secret = "ldfjklds" //盐值

func main() {

	oPassword := "123456" //旧密码
	h := md5.New()
	h.Write([]byte(secret))
	fmt.Println(hex.EncodeToString(h.Sum([]byte(oPassword)))) //得到16进制的编码

	//字符串遍历
	s1 := "你好啊，世界"
	for k, v := range s1 {
		fmt.Println(k, v, string(v))
	}

	//字符串转换
	s2 := []byte(s1) //18个字节，一个汉字3个字节
	fmt.Println(s2, len(s2))
	fmt.Println(string(s2))

}
