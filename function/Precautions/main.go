package main

import (
	"fmt"
)

func getSum(n1 int, n2 int) int  {
	return n1 + n2
}

type myFunc func(int, int) int
func myFun(funvar myFunc, num1 int, num2 int ) int {
	return funvar(num1, num2)
}

//支持对函数返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int)  {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//可变参数的使用
func sum(n1 int, args... int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main()  {

	a := getSum
	fmt.Printf("a的类型是%T\ngetSum的类型是%T\n", a, getSum)

	res := a(20, 30)
	fmt.Println("res = ", res)

	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2 = ", res2)

	type myInt int

	var number myInt
	var number1 int
	number = 40
	number1 = int(number)
	fmt.Println("number = ", number, "number1 = ", number1)


	a1, b1 := getSumAndSub(2, 3)
	fmt.Printf("a1 = %v b1 = %v\n", a1, b1)


	res4 := sum(10, -1, 20, 111, 30)
	fmt.Println("res4 = ", res4)

}