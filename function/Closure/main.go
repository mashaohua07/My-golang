package main

import (
	"fmt"
	"strings"
)

//累加器
func AddUpper() func (int) int  {
	var n int = 10
	var str string = "hello"
	return func(x int) int {
		n = n + x
		str += string(36)
		fmt.Println("str = ", str)
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	
	return func(name string) string {
		if !strings.HasSuffix(name, suffix)  {
			return name + suffix
		}
		return name
	}
}

func makeSuffox2(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func main()  {

	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后的返回:", f2("kill"))
	fmt.Println("文件名处理后的返回:", f2("kill.jpg"))

	fmt.Println("文件名处理后的返回:", makeSuffox2(".jpg", "kill"))
	fmt.Println("文件名处理后的返回:", makeSuffox2(".jpg", "kill.jpg"))
}
