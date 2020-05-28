package main

import (
	"fmt"
	"function/utils"
)

var age = test()

func test() int {
	fmt.Println("test()")
	return 90
}

//init函数，通常可以在init函数中完成初始化的工作
func init()  {
	fmt.Println("init()...")
}

func main()  {
	fmt.Println("main()...age", age)
	fmt.Println("Age = ", utils.Age, "Name = ", utils.Name)
}