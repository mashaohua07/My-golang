package main

import "fmt"

type Stu struct {
	Name string
	Age int
}

func main()  {
	su := &Stu{"doin", 50}
	fmt.Println(*su)
}