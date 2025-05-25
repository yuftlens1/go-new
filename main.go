// 强制推荐使用 Go Modules（所有新项目）  ;     切换到你的 项目根目录（即包含 main.go 和 db/ 文件夹的目录）   cd ~/go/src/study && go mod init github.com/yuftlens1/go-new && cat go.mod  #!!!
package main //#声明man包，声明是主函数执行。所有的go语言都是从main这个包开始执行！

import (
	"fmt" //导入现有的输入输出的框架

	"github.com/yuftlens1/go-new/db" //导入自己写的db包      对应上面的go mod init

	"github.com/yuftlens1/go-new/iffor"
)

func main() { //#func代表一个函数
	fmt.Println("加油")
	fmtEnvv := fmt.Sprintf("Environment: %s, Version: %d", "百分号S的参数;后面十进制表示", 0x12)
	fmt.Println(fmtEnvv)

	fmt.Println("Todo list", db.Todos) //调用db包里的变量 Todos, 因为大写开头所以可以被其他包调用

	iffor.Runif()                   //调用iffor包里的函数 Runif, 因为大写开头所以可以被其他包调用  ;         格式:             包名.函数名()

	iffor.Runfor()

}

/*
fmt.Sprintf  是最常用的格式化函数，它根据格式化字符串和参数生成一个格式化后的字符串。
fmt.Sprint   用于将多个参数连接成一个字符串，但不支持格式化占位符。
fmt.Sscanf   fmt.Sscanf 用于从字符串中解析数据，与 fmt.Sprint 和 fmt.Sprintf 的功能不同。
*/
