package main //必须有一个main包

import "fmt" //导入包 必须要使用

func main() {
	//变量，程序运行期间可以改变的量

	//1.声明格式  var 变量名 类型，变量声明了必须要使用
	//2.只是声明没有初始化的变量，默认值为0
	//3.同一个{}声明的变量名是唯一的
	var a int
	fmt.Println("a  = ", a)

	//4.可以同时声明多个变量
	var b, c, d int
	b = 3
	c = 6 //变量的赋值
	d = a + b + c

	fmt.Println(d)

	//5.变量的初始化，声明变量，同事赋值
	var e int = 9 //初始化，声明变量时，同时赋值（一步到位）
	e = 3         //赋值:先声明，后赋值
	fmt.Println("e =", e)

	//6.自动推导类型，必须初始化，通过初始化的值确定类型
	f := 30
	//%T打印变量所属的类型
	fmt.Println("f type is %T\n", f)

}
