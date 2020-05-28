package main

import (
	"fmt"
	"strconv"
	_ "unsafe"
)

func main()  {

	var num1 int = 99
	var num2 float64 = 23.111
	var b bool =true
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)


	var num3 int = 99
	var num4 float64 = 23.111
	var c bool =true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 6, 64)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = strconv.FormatBool(c)
	fmt.Printf("str type %T str=%v\n", str, str)

	var num5 int64 = 123
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str=%q\n", str, str)
}