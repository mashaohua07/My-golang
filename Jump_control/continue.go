package main

import "fmt"

func main() {

	//continue案例
	//for i := 0; i < 4; i++ {
	//	for j := 0; j < 10; j++ {
	//		if j == 2 {
	//			continue
	//		}
	//		fmt.Println("j = ", j)
	//	}
	//}


	//continue实现，打印1-100之内的奇数【要求使用for循环+continue】
	//for i := 1; i <= 100; i++ {
	//	if i % 2 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}


	//从键盘读入个数不确定的整数，并判断读入的正数和负数的个数，输入为0时结束程序
	var positiveCount int	//正数个数
	var negativeCount int	//负数个数
	var num int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("正数的个数是%v, 负数的个数是%v\n", positiveCount, negativeCount)
}