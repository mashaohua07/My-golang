package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	myByte := [26]byte{}
	for i := 0; i < 26; i++ {
		myByte[i] = 'A' + byte(i)
	}
	fmt.Printf("myByte = %c\n", myByte)


	myInt := [...]int{1, -20, 56, 300, -54, -555, 0}
	myVal := myInt[0]
	myValindex := 0
	for i := 1; i < len(myInt); i++ {
		if myVal < myInt[i] {
			myVal = myInt[i]
			myValindex = i
		}
	}
	fmt.Printf("myVal = %v, myValindex = %v\n", myVal, myValindex)


	//随机生成5个数，并将其反转打印
	intArr := [5]int{}
	len := len(intArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)
	//反转打印
	temp := 0
	for i := 0; i < len / 2; i++ {
		temp = intArr[len -1 -i]
		intArr[len -1 -i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println(intArr)

}