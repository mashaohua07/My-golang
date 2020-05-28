package main

import (
	"fmt"
	"strconv"
	"time"
)

func test()  {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main()  {
	//在执行test()前，先获取当前Unix时间戳
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("执行程序共花费时间为%v\n", end-start)
}
