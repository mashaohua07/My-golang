package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	//time.Now().Unix() 从1970:01:01 的0时0分0秒到现在的秒数
	//rand.Seed(time.Now().Unix())
	//n := rand.Intn(100)
	//fmt.Println(n)

	//随机生成1-100的一个数，直到生成99这个数，看看一共用了几次
	var count int = 0
	rand.Seed(time.Now().UnixNano())
	for  {
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break //表示跳出for循环
		}
	}
	fmt.Println("生成 99 一共用了", count, "次")


	//演示指定标签的形式来使用break
	lable1:
	for i := 0; i < 4; i++ {
		//lable2:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break
				break lable1
			}
			fmt.Println("j = ", j)
		}
	}


}