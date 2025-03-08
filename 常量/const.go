package main

import (
	"fmt"
	"reflect"
)

const c1, c2, c3 = 1, 2, 3    //var 是变量，const是常量

func main() {
	const c1, c2, c3, a  = 11, 22, 33 ,true

	fmt.Println(c1,a,)
	fmt.Println(reflect.TypeOf(c1),reflect.TypeOf(a))
}