package main

import (
	"fmt"
)

func main()  {
	local := make(map[string]string)
	local["string"] = "var"
	local["int"] = "18"
	local["del"] = "error"
	fmt.Println(local)

	delete(local, "int")
	fmt.Println(local)

	delete(local, "install")
	fmt.Println(local)

	//如果希望一次性删除所有的key
	//1.遍历所有的key，然后逐一删除
	//2.直接make一个新的空间
	//local = make(map[string]string)
	//fmt.Println(local)

	//map中的查找
	val, ok := local["del"]
	if ok {
		fmt.Printf("找到del 值为%v\n", val)
	}else {
		fmt.Println("没有找到del")
	}
}
