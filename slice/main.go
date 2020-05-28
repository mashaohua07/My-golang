package main

import (
	"fmt"
)

func main()  {
	intArr := [...]int{11, 22, 33, 44, 55}
	slice := intArr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println()

	slicenew := make([]int,5,10)
	slicenew[0] = 10
	slicenew[3] = 20
	fmt.Println(slicenew)
	fmt.Println(len(slicenew))
	fmt.Println(cap(slicenew))

	fmt.Println()

	sliceold := []string{"etcd", "APIServer", "scheduler", "docker", "kubelet", "contreller", "proxy"}
	fmt.Println(sliceold)
	fmt.Println(len(sliceold))
	fmt.Println(cap(sliceold))

	//用append内置函数，可以对切片进行动态追加
	sliceold = append(sliceold, "pod", "node")
	fmt.Println(sliceold)
	//通过append将切片自我增长
	sliceold = append(sliceold, sliceold...)
	fmt.Println(sliceold)

	//切片的拷贝操作
	slicepod := []int{10, 20, 30, 40, 50}
	slicepod1 :=make([]int, 10)
	copy(slicepod1, slicepod)
	fmt.Println(slicepod)
	fmt.Println(slicepod1)
}
