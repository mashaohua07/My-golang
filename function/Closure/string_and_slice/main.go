package main

import (
	"fmt"
)

func main()  {
	str := "mashaohua07@sina.com"
	slice := str[:9]
	fmt.Println(slice)
	//将str中的07转换为11，因为string不能直接修改，需要转化为切片
	arr1 := []byte(str)
	arr1[9] = '1'
	arr1[10] = '1'
	fmt.Printf("%s\n", arr1)
	str = string(arr1)
	fmt.Println(str)

	//使用[]rune替换中文
	arr2 := []rune(str)
	arr2[0] = '中'
	str = string(arr2)
	fmt.Println(str)
}
