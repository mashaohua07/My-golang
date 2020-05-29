package main

import (
	"fmt"
)

func main()  {

	//for-range 遍历数组
	helm := [...]string{"kubernetes", "istio", "docker", "harbor"}

	for index, value := range helm {
		fmt.Printf("index = %v, value = %v\n", index, value)
	}

	for _, value := range helm{
		fmt.Printf("value = %v\n", value)
	}
}
