package main

import (
	"fmt"
	_ "fmt"
)

type AInterface interface {
	Say()
}

type BInterface interface {
	Run()
}

type CInterface interface {
	AInterface
	BInterface
	Len()
}

type Stu struct {

}

func (stu Stu) Say() {

}

func (stu Stu) Sun() {

}

func (stu Stu) Len() {

}

type T interface {

}

func main()  {
	var stu Stu
	var a AInterface = stu
	a.Say()

	var t T = stu
	fmt.Println(t)
}
