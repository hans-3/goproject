package main

import (
	"fmt"
)
func main() {
	//1.声明变量  没有初始化零值（）为false
	var a bool
	a = true
	fmt.Println("a0 = ", a)

	a = true
	fmt.Println("a = ", a)
	//自动推导类型
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c = ", c)
}