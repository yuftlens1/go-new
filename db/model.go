package db                                  //一个目录是一个包,包里的各个go文件的 package 尽量和目录名字一致   有助于代码的组织和导入

type Todo struct { //自定义结构体   自定义数据类型  非常重要的复合数据类型
	/*
		结构体标签(Struct Tag)：
			`json:"id"` 是一个结构体标签 ;      标签提供了关于字段的元信息
			json:"id" 表示当这个结构体被序列化为 JSON 时：
				字段 ID 在 JSON 中会显示为小写的 "id"
				同样地，从 JSON 反序列化时，"id" 会被映射到 ID 字段
	*/
	ID string `json:"id"` //ID string 定义了一个名为 ID 的字符串类型字段      `json:"id"` 可以不写

}

var Todos []Todo                    //定义全局变量名字Todos , 类型是 []Todo     ;   即 Todo 结构体的切片（动态数组）
/*
这个变量可以被同一个包中的所有文件访问
因为是首字母大写，所以也可以被其他包导入访问（公开的）
*/

var todos []Todo //只能被 db 目录下的代码引用
