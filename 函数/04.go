package main
import "fmt"

var (
	name string = "haha"
)

func printAd()  {                //函数:一个函数最好做一个功能
	fmt.Println("程序主体")
	fmt.Println("第一步：学习go语言")
	fmt.Println("第二步：用go语言开发")
	id := name                   //全局变量赋予到局部变量
	fmt.Println("函数Ad里调用",id)
}



func main() {                    //go语言执行体 main 函数
	fmt.Println("程序开始运行")
	printAd()                    //调用函数
	fmt.Println("程序运行完成")
	//fmt.Println("main函数里调用",id)    //局部变量不能跨函数
}