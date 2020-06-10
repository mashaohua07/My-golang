package main

import "fmt"

func main()  {
	var Inter interface{}
	var b float64 = 20.1
	Inter = b

	if y, kil := Inter.(float64); kil {
		fmt.Println("success")
		fmt.Printf("%T, %v\n", y, y)
	} else {
		fmt.Println("failure")
	}
	fmt.Println("interface")

}
