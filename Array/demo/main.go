package main

import (
	"fmt"
)

func main()  {
	//数组是值类型
	hen := [...]float64{2, 3, 4, 5.5, 6, 7, 8}
	total := 0.0
	for i := 0; i < len(hen); i++ {
		total += hen[i]
	}

	fmt.Printf("helm = %v\n", total)

	average := fmt.Sprintf("%.2f", total / float64(len(hen)))
	fmt.Printf("average = %v\n", average)

	fmt.Printf("hen的地址: %p, hen[0]的地址: %v, hen[1]的地址: %v\n", &hen, &hen[0], &hen[1]) //int占8个字节


	// 从终端循环输入5个成绩，保存到float64数组，并输出
	score := [5]float64{}

	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值\n", i+1)
		fmt.Scanln(&score[i])
	}

	//变量数组打印
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}
	fmt.Printf("srcre = %v\n", score)


	//可以指定元素值对应的下标
	strArr := [...]string{2:"stop", 0:"start", 1:"restart", 3:"enable"}
	fmt.Printf("strArr的值: %v\n", strArr)

}