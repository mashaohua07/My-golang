package main

import (
	"fmt"
)

var (
	//fun1就是一个全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main()  {

	//使用匿名函数的方式求俩数的和
	res := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res = ", res)


	//将匿名函数赋值给一个变量来使用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res1 := a(50, 30)
	fmt.Println("res1 = ", res1)


	//全局匿名函数的使用
	res2 := Fun1(4, 4)
	fmt.Println("res2 =", res2)

}
