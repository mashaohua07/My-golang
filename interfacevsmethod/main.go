package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}

type Bird interface {
	Flying()
}

type Little struct {
	Monkey
}

func (mon Monkey) climb() {
	fmt.Println("climbing")
}

func (fl *Little) Flying() {
	fmt.Println("Flying")
}

func main()  {

	monkey := Little{
		Monkey{
			"kong",
		},
	}

	monkey.climb()
	monkey.Flying()
}
