package main

import (
	"fmt"
	"sort"
)

type Stu struct {
	Name string
	Age int
	mt string
}

func main()  {
	mon := make(map[int]int)
	mon[60] = 20
	mon[50] = 155
	mon[40] = 5
	mon[70] = 55

	var keys []int
	for key, _ := range mon {
		keys = append(keys, key)
	}

sort.Ints(keys)
fmt.Println(mon)

	for _, k := range keys {
		fmt.Printf("mon[%v]=%v\n", k, mon[k])
	}

	fond := make(map[string]Stu, 5)
	n := Stu{"tom", 18, "ommint"}
	s := Stu{"jak", 17, "openist"}
	fond["no1"] = n
	fond["no2"] = s
	fmt.Println(fond)

	for k, v := range fond {
		fmt.Printf("编号=%v\n", k)
		fmt.Printf("Name=%v\n", v.Name)
		fmt.Printf("Age=%v\n", v.Age)
		fmt.Printf("Mt=%v\n", v.mt)
	}
}