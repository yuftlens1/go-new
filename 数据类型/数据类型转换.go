package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	aa := 1
	fmt.Println("变量aa的数据类型是:",reflect.TypeOf(aa))

	var bb int32 = 11
	fmt.Println("变量bb的数据类型是:",reflect.TypeOf(bb))

	fmt.Println("数据类型int的取值范围:",math.MaxInt,math.MinInt)           //和int64的一样，因为我的操作系统是64位的
	fmt.Println("数据类型int16的取值范围:",math.MaxInt16,math.MinInt16)
	fmt.Println("数据类型int32的取值范围:",math.MaxInt32,math.MinInt32)
	fmt.Println("数据类型int64的取值范围:",math.MaxInt64,math.MinInt64)    
}

//int的取值范围和操作系统有关系. 64位操作系统是int64,32位操作系统是int32.