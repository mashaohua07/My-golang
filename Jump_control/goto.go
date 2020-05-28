package main

import "fmt"

func main()  {

	//goto的使用
	var n int =  5
	fmt.Println("kill 3")
	if n < 10 {
		goto label
	}
	fmt.Println("kill 4")
	fmt.Println("kill 5")
	label:
	fmt.Println("kill 6")
	fmt.Println("kill 7")
}