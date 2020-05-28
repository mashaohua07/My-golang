package main

import "fmt"

func main()  {

	var age int = 10

	switch {
		case age == 10 :
			fmt.Println("age == 10")
		case age == 20:
			fmt.Println("age == 20")
		default:
			fmt.Println("没有匹配")
	}

	//var n1 int64 = 55
	//var n2 int64 = 66
	//
	//switch n1 {
	//	case n2, 10, 50 :
	//		fmt.Println("n1 == ", n1)
	//	case 55 :
	//		fmt.Println("n1 == ", 55)
	//}

	//switch 的穿透 fallthrought
	var num int = 10
	switch num {
		case 10:
			fmt.Println("ok1")
			fallthrough	//默认穿透一层
		case 20:
			fmt.Println("ok2")
		case 30:
			fmt.Println("ok3")
		default:
			fmt.Println("没有匹配")

		
		
	}

}


