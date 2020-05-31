package main

import (
	"fmt"
)

func BuddleSort(arr *[]int)  {
	fmt.Printf("%v\n", *arr)
	temp := 0
	for j := 0; j < len(*arr) - 1; j++ {
		for i := 0; i < len(*arr) - 1 - j; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
		}
	}
	fmt.Printf("%v\n", *arr)
}

func main()  {
	arr := []int{23, 10, 50, 66, 2}
	BuddleSort(&arr)
	fmt.Printf("%v\n", arr)
}
