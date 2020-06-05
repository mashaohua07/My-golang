package main

import (
	"fmt"
)

type Persion struct {
	Name string
	Age int
	Scoure [5]float64
	ptr *int
	slive []int
	map1 map[string]string
}

func main()  {
	var p1 Persion

	p1.slive = make([]int, 5)
	p1.slive[0] = 100

	p1.ptr = new(int)

	p1.map1 = make(map[string]string)
	p1.map1["int"] = "var"

	fmt.Println(p1)

	p2 := new(Persion)
	(*p2).Name = "python"
	p2.Age = 50
	p2.ptr = &(p2.Age)

	fmt.Println(p2)
	fmt.Println(*(p2.ptr))

}
