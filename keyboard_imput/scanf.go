package main

import "fmt"

func main()  {

	var name string
	var age int
	var sal float32
	var isPass bool
	fmt.Println("plase name")
	fmt.Scanf("%s %d %f %t" , &name, &age, &sal, &isPass)
	fmt.Printf("name %v \n age %v \n sal %v \n isPass %v \n", name, age, sal, isPass)
}
