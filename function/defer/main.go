package main

import (
	"fmt"
)

//当执行到defer时，暂时不执行，会将defer后面的语句压入独立的栈
//当函数执行完事后，再从defer栈，按照先进后出，执行
func sum(n1 int, n2 int) int {
	defer fmt.Println("n1 = ", n1)
	defer fmt.Println("n2 = ", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("res = ", res)
	return res
}

func main()  {
	res := sum(10, 20)
	fmt.Println("res1 = ", res)
}