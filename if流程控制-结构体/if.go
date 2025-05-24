package main

import "fmt"

func main() {

	jiage := 100

	var tianqi string
	tianqi = "qiangtian"

	if tianqi != "qiangtian" { //!=是不等于;  ==是等于
		fmt.Printf("今天天气是%v，今天雨伞的价格是%vyuan", tianqi, jiage)
	} else {
		fmt.Printf("今天雨伞的价格是%+v元\n", jiage+10)
	}

	helper() //调用一个包里的其他函数  ; 一个包里可以直接调用其他函数，不需要写import ;小写的变量 常量 和函数  可以在一个包里直接调用
	runtestclass()
}

//Printf  格式化输出-变量;   %v值的默认格式;万能匹配啊           %+v 类似%v，但输出结构体字段名 得有变量定义那边配合应该才能用起来
//%s 格式化输出字符串;     %d 十进制
