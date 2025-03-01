package main //#声明man包，声明是主函数执行。所有的go语言都是从main这个包开始执行！

import "fmt" //#导入现有的输入输出的框架

func main() { //#func代表一个函数
	fmt.Println("加油aa")
	fmtEnvv := fmt.Sprintf("Environment: %s, Version: %d", "百分号S的参数;后面十进制表示", 0x12)
	fmt.Println(fmtEnvv)
}

/*
fmt.Sprintf  是最常用的格式化函数，它根据格式化字符串和参数生成一个格式化后的字符串。
fmt.Sprint   用于将多个参数连接成一个字符串，但不支持格式化占位符。
fmt.Sscanf   fmt.Sscanf 用于从字符串中解析数据，与 fmt.Sprint 和 fmt.Sprintf 的功能不同。
*/
