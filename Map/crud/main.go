package main

import (
	"fmt"
)

func main()  {
	local := make(map[string]string)
	local["string"] = "var"
	local["int"] = "18"
	local["del"] = "error"
	fmt.Println(local)

	//delete(local, "int")
	//fmt.Println(local)
	//
	//delete(local, "install")
	//fmt.Println(local)

	//如果希望一次性删除所有的key
	//1.遍历所有的key，然后逐一删除
	//2.直接make一个新的空间
	//local = make(map[string]string)
	//fmt.Println(local)

	//map中的查找
	val, ok := local["del"]
	if ok {
		fmt.Printf("找到del 值为%v\n", val)
	}else {
		fmt.Println("没有找到del")
	}

	//map的遍历
	for k, v := range local {
 		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	newmap := make(map[string]map[string]string)
	newmap["status"] = make(map[string]string)
	newmap["status"]["top"] = "api"
	newmap["status"]["pod"] = "sch"
	newmap["status"]["node"] = "lete"
	newmap["start"] = make(map[string]string)
	newmap["start"]["top"] = "server"
	newmap["start"]["pod"] = "service"
	newmap["start"]["node"] = "system"
	fmt.Println(newmap)
	for k, v := range newmap {
		fmt.Println("k=", k)
		for k1, v1 := range v {
			fmt.Printf("\t k1=%v, v1=%v\n", k1, v1)
		}
	}
}
