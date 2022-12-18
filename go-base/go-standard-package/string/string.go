package main

import (
	"fmt"
	"strings"
)

//包简单的用法
func main() {
	/**
	--------->字符串长度<------------
	*/
	fmt.Println(len(""))
	fmt.Println(len("123abc!@#"))
	fmt.Println(len("我是中国人"))

	/**
	--------->字串包含关系<------------
	*/
	fmt.Println("--------->字串包含关系<------------")
	//substr在str中出现的次数
	fmt.Println(strings.Count("ababa", "ab"))
	//判断子串是否在给定的字串中，不存在就返回false
	fmt.Println(strings.Contains("0080", "800"))
	//判断子串中任何一个字符是否在s中
	fmt.Println(strings.ContainsAny("0080", "001"))
	fmt.Println(strings.ContainsAny("中国", "人中龙"))
	//判断
	strr := "a中国"
	strrr := []rune("a中国")
	fmt.Printf("type : %T, val : %v\n", strr, strr)
	fmt.Printf("type : %T, val : %v\n", strrr, strrr)
	fmt.Println(strings.ContainsRune("中国", 22269))
	//返回子串第一次出现的下标值
	fmt.Println(strings.Index("abcdefggff", "f"))
	//返回子串最后一次出现的下标值
	fmt.Println(strings.LastIndex("abcdefggff", "f"))
	fmt.Println(strings.LastIndex("abcdefggff", "h"))

	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))

	/**
	--------->字符串分割<------------
	*/
	fmt.Println("--------->字符串分割<------------")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a"))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 3))
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d,ef,dd", ",", 4)) //按指定分隔符分割成几份
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	/**
	--------->字符串替换<------------
	*/
	fmt.Println("--------->字符串替换<------------")
	//字符串重复，1是原字符串
	fmt.Println(strings.Repeat("中国", 3))
	//从左到右替换两次，不是只替换替换第二次
	fmt.Println(strings.Replace("我是中国人是不是？", "是", "不是", 2))
	fmt.Println(strings.ReplaceAll("我是中国人是不是？", "是", "不是"))

	str := "12345"
	str1 := []byte("12345")
	fmt.Println(str[0] == str1[0])
	fmt.Printf("type : %T, val : %v\n", str[0], str[0])
	fmt.Printf("type : %T, val : %v\n", str1[0], str1[0])

	str2 := "abcdef"
	str3 := []byte("abcdef")
	fmt.Printf("type : %T, val : %v\n", str2[0], str2[0])
	fmt.Printf("type : %T, val : %v\n", str3[0], str3[0])

	str4 := "!!@*^*&"
	str5 := []rune("!!@*^*&")
	str6 := str5[:]
	fmt.Printf("type : %T, val : %v\n", str4[0], str4[0], str4[0])
	fmt.Printf("type : %T, val : %v\n", str5[0], str5[0], str4[0])
	fmt.Println(str6, string(str6))

}
