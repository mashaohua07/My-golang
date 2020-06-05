package main

import (
	"fmt"
	"method/factory/model"
)

func main()  {
	//stu := model.Student{
	//	Name: "node",
	//	Score: 30.2,
	//}
	stu := model.NewStudent("node", 30.2)

	fmt.Println(*stu)
	fmt.Println("Name = ", stu.Name, "Score = ", stu.GetStudent())
}