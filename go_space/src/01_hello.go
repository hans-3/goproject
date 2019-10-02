package main //每个程序必须有一个main包

import "fmt"

func main() {
	fmt.Println("life is shit")
	fmt.Println("hello")
	a, b := 123, 321
	c, d := "dsds", 5.5
	fmt.Print("你好，世界！")             //输出控制台
	fmt.Println(a)                  //输出控制台后自动换行
	fmt.Printf("a=%d\n", a)         //格式化输出
	a, b = b, a                     //交换值
	fmt.Printf("a=%d，b=%d\n", a, b) //%d输出整型
	fmt.Printf("d=%f\n", d)         //%f输出浮点型
	fmt.Printf("a type is %T\n", a) //%T输出类型
	fmt.Printf("a 为:%v\n", a)       //%v自动匹配输出
	fmt.Printf("a 为:%c\n", a)       //%c输出字符型（按asc码输出）
	fmt.Printf("c 为:%s\n", c)       //%s输出字符串
	var bo bool
	fmt.Printf("bo 布尔为：%t\n", bo) //%t输出布尔类型

	fmt.Println("请输入：")
	fmt.Scan(&c)
	fmt.Printf("你输入的是：%v\n", c)

}
