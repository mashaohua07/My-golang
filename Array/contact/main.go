package main

import (
	"fmt"
)

func main()  {

	//创建一个byte类型的26个元素的数组，分别设置'A'-'Z'.使用for循环访问说有元素并打印出来.提示：字符数据运算 'A'+ 1 -> 'B'

	//思路
	//1. 声明一个数组
	//2. 使用for循环，利用字符可以进行运算的特点来赋值 'A'+ 1 -> 'B'
	myChars := [26]byte{}
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}
	fmt.Printf("myChars = %c\n", myChars)


	//请求出一个数组的最大值，并得到对应的下标

	//思路
	//1. 声明一个数组
	//2. 假定第一个数是最大的，下标就是0
	//3. 然后从第二个数循环比较，如果发现有更大的，则交换
	myInt := [...]int{1, -1, 10, 30, 110, -50, 0, 9000}
	myVal := myInt[0]
	myValIndex := 0
	for i := 0; i < len(myInt); i++ {
		if myVal < myInt[i] {
			myVal = myInt[i]
			myValIndex = i
		}
	}
	fmt.Printf("myVal = %v, myValIndex = %v\n", myVal, myValIndex)
}