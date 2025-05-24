package main

import "fmt"

func helper() {
	fmt.Println("我是 utils.go 里的 helper 函数")
}

//Printf  格式化输出-变量;   %v值的默认格式;万能匹配啊           %+v 类似%v，但输出结构体字段名 得有变量定义那边配合应该才能用起来
//%s 格式化输出字符串;     %d 十进制
