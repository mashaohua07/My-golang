package main

import "fmt"

func main()  {
	arr := [4][6]int{}
	arr[0][2] = 1
	for _, value := range arr{
		fmt.Println()
		for _, r := range value{
			fmt.Print(r, " ")
		}
	}
	fmt.Println()
}