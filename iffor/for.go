package iffor

import (
	"fmt"
	"time"
)

func Runfor()  {

	
	count := 0         //定义变量;  := 自适应匹配数据类型
	jishu := 0         //定义变量;  := 自适应匹配数据类型

	for i := 0; i < 101; i++ {                 // i等于0 ; i不能大于101 ;     i++的意思就 +1   (i = i + 1 的简写)
		if i%2 == 0{                           // 只有一段if  没有 else if 或 else  也能跑 
			fmt.Println("发现一个偶数:", i)
			count++
		} else {
			fmt.Println("发现一个奇数:", i)
			jishu++
		}

	}

	fmt.Println("偶数共有:", count)
	fmt.Println("奇数共有:", jishu)


	for  {                                //无限循环
		time.Sleep(time.Second * 2)       //每2秒执行


		date := time.Now()
		fmt.Println(time.Now())                              //原始格式输出当前时间      date := time.Now() 赋予变量里 
		fmt.Println(date.Format("2006-01-02 15:04:05"))      //格式化输出时间 年月日 时分秒
		fmt.Println(date.Format("2006-01-02"))               //格式化输出时间 年月日
		fmt.Println(date.Format("2006-01-02 15:04"))         //格式化输出时间 年月日 时分
	}

}