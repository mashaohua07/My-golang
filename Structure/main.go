package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
	color string
	arr [3]int
}

func main()  {

	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.color = "白色"
	fmt.Println(cat1)

	fmt.Println(cat1.Name)
	fmt.Println(cat1.Age)
	fmt.Println(cat1.color)
	fmt.Println(cat1.arr)
}
