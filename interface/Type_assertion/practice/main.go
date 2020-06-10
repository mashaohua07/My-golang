package main

import "fmt"

func Practice(it... interface{})  {

	for index, x := range it {
		switch x.(type) {
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index, x)
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, x)
		case int:
			fmt.Printf("第%v个参数是int类型，值是%v\n", index, x)
		case byte:
			fmt.Printf("第%v个参数是byte类型，值是%v\n", index, x)
		default:
			fmt.Println("第%v个参数是不确定类型，值是%v\n", index, x)
		}
	}
}

func main()  {
	var n1 float32 = 1.1
	var n2 float64 = 2.2
	var n3 bool = true
	var n4 string = "morker"
	var n5 int = 55
	var n6 byte = 15
	address := "杭州"
	Practice(n1, n2, n3, n4, n5, n6, address)
}