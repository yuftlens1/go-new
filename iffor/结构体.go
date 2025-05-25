package iffor

import (
	"fmt"
	"reflect"
)

//相当于自定义数据类型;类型合集?    非常重要的复合数据类型

type testclass struct { //定义新的数据的类型 testclass
	Age  int
	Name string
}

var b = testclass{Name: "yufthaha"}

// b = testclass{Name: "yuft"}       // 会报错，要么上面定义cc的时候直接赋值    ；    要么将赋值移到函数体内  ;

func runtestclass() {
	var a testclass
	a = testclass{Age: 29}

	fmt.Println(a.Age, b.Name)
	fmt.Printf("我的年龄是%v。我的名字是%v,谢谢\n", a.Age, b.Name)   //格式化输出 %v 是通用输出符号, 输出相应值的默认格式.
	fmt.Printf("我的年龄是%+v。我的名字是%+v,谢谢\n", a.Age, b.Name) //更详细的通用占位符 %+v - 对结构体会显示字段名
	fmt.Printf("我的年龄是%#v。我的名字是%#v,谢谢\n", a.Age, b.Name) //%#v 输出值的Go语法表示

	fmt.Println(reflect.TypeOf(a))
}
