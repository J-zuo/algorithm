package main

import (
	"fmt"
)

//探究为什么map遍历时是无序的，怎样实现有序遍历？

func printMap(cityMap map[string]string) {
	//你会发现，每次遍历,key输出的顺序是不一定的，是无序的，什么原因？
	for k, v := range cityMap {
		fmt.Printf("key:%v, value:%v\n", k, v)
	}

}

func addElenment(cityMap map[string]string, ele string, elev string) {
	cityMap[ele] = elev
}

func main() {
	//简单的示例
	//定义
	m2 := make(map[string]string, 13)
	fmt.Println(m2, len(m2))

	//赋值
	m2["张三"] = "老师"
	m2["李四"] = "军人"
	//从结果可以看出，map的key是不能重复的

	//遍历
	printMap(m2)

	fmt.Println("---------")
	//修改（新增和修改）
	addElenment(m2, "王武", "护士")   //新增
	addElenment(m2, "张三", "司机")   //修改原key的值
	addElenment(m2, "张三1", "司机")  //修改原key的值
	addElenment(m2, "张三2", "司机")  //修改原key的值
	addElenment(m2, "张三3", "司机")  //修改原key的值
	addElenment(m2, "张三4", "司机")  //修改原key的值
	addElenment(m2, "张三5", "司机")  //修改原key的值
	addElenment(m2, "张三6", "司机")  //修改原key的值
	addElenment(m2, "张三7", "司机")  //修改原key的值
	addElenment(m2, "张三11", "司机") //修改原key的值
	addElenment(m2, "张三8", "司机")  //修改原key的值
	addElenment(m2, "张三9", "司机")  //修改原key的值
	addElenment(m2, "张三10", "司机") //修改原key的值
	fmt.Println(m2, len(m2))

	printMap(m2)
	fmt.Println("=-------")
	//删除
	delete(m2, "李四") //如果delete删除一个不存在的，对原map没有影响

	printMap(m2)

}
