package main

import (
	"fmt"
	"project/Array"
)

func main()  {

	//数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，就不能动态变化
	arr01 := [3]int{}
	arr01[0] = 5
	arr01[1] = 6
	arr01[2] = 8
	//arr01[3] = 10		//error，一旦定义长度，就不能动态变化
	fmt.Printf("arr01 = %d\n", arr01)

	arr02 := [3]string{"master", "node"}	//string类型，空值默认为""
	fmt.Printf("arr02 = %q\n", arr02)

	arr03 := [3]bool{false, true}	//bool类型，空值默认为false
	fmt.Printf("arr03 = %v\n", arr03)


	//Go的数组属值类型，在默认情况下是值传递，因此会进行值拷贝，数组间不会相互影响
	arr := [3]int{10, 20, 30}
	Array.Test(arr)
	fmt.Printf("arr = %v\n", arr)

	//如果想在其他函数中，去修改原来的数组，可以使用指针传递
	arr1 := [3]int{40, 50, 60}
	Array.Test1(&arr1)
	fmt.Printf("arr1 = %v\n", arr1)

}