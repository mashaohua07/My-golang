package main

import (
	"fmt"
	"strconv"
)

func main()  {

	var str string = "true"
	var b bool
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "123456789"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)


	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)

	//注意
	var str4 string = "hello"
	var n5 int64
	n5, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n5 type %T n5=%v\n", n5, n5)
}
