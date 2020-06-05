package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
}

func (stu *Student) String() string  {
	str := fmt.Sprintf("Name = [%v], Age = [%v]", stu.Name, stu.Age)
	return str
}

func main()  {

	stu := Student{
		Name : "look",
		Age : 20,
	}
	fmt.Println(&stu)		//对于定于了String()方法的类型，默认输出的时候会调用该方法，实现字符串的打印
	fmt.Println((&stu).String())

}
