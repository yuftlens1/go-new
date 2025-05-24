package main //#声明man包，声明是主函数执行。所有的go语言都是从main这个包开始执行！
import (
	"fmt" //#导入现有的输入输出的框架
	"reflect"
)

/*
go是静态语言。静态语言和动态语言的变量相比差异很大

go语言变量定义：
1.变量不赋值默认是0 或者 空字符
2.变量必须有类型，而且变量定义的类型和其数据类型必须一致
3.类型定下来后不能改变
4.局部变量定义了变量必须使用，全局变量不是必须使用
5.全局变量不能用简洁变量， := 定义
6.变量名不能冲突. 全局变量和局部变量可以名称冲突，局部变量优先级高
*/

//定义变量的方式: 类型在最后 ;   变量类型 和 数据类型 是同一个概念，它们都用来描述变量可以存储的数据的种类和范围。

var name string = "haha" //写到func main是全局变量，所以它的作用域是整个包。
var aa string = "aa"

var cc string

// cc = "testcc"   会报错，要么上面定义cc的时候直接赋值    ；    要么将赋值移到函数体内（推荐）

// age := 12   #全局变量不能用:= 定义, 必须写var   包级变量

var ( //单变量
	namea      = "woqu"
	num        = 19
	ok    bool = true
	okk   bool //布尔默认false
)

var user1, user2, user3 = "bobby1", 1, "bobby3"

func main() {
	age := 12 //局部变量，必须使用
	//num := 13
	fmt.Println(name, age, namea, ok, okk, num)
	fmt.Println(user1, user2, user3)
	fmt.Println("ok 该变量的类型是:", reflect.TypeOf(ok)) //输出变量的数据类型
}

/*
##基本数据类型：
bool：布尔类型，值为 true 或 false
int：有符号整数，位数根据系统架构（通常是 32 位或 64 位）
int8：有符号 8 位整数
int16：有符号 16 位整数
int32：有符号 32 位整数
int64：有符号 64 位整数
uint：无符号整数，位数根据系统架构（通常是 32 位或 64 位）
uint8：无符号 8 位整数
uint16：无符号 16 位整数
uint32：无符号 32 位整数
uint64：无符号 64 位整数
float32：32 位浮点数
float64：64 位浮点数
complex64：64 位复数，由两个 float32 组成
complex128：128 位复数，由两个 float64 组成
string：字符串类型，由字节序列组成
byte：等同于 uint8，用于表示字节
rune：等同于 int32，用于表示 Unicode 码点

##复合数据类型：
数组：[3]int 表示长度为 3 的整数数组
切片：[]int 或 []string 等，动态数组
结构体：struct，用于表示记录（row）类型
映射：map[string]int 等，键值对集合
函数：func(x int, y int) int 等，表示函数类型
通道：chan int 等，用于并发通信

##指针类型：
*int：指向整数的指针
*MyStruct：指向自定义结构体 MyStruct 的指针

##接口类型：
interface{}：空接口，可以存储任意类型的值
自定义接口：如 type MyInterface interface { Method1(); Method2() }
*/
