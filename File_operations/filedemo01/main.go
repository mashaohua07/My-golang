package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	file := "../test"
	by, err := ioutil.ReadFile(file)	//ioutil.ReadFile一次性将文件读取到位，只适合文件不太大的操作，否则效率低下
	if err != nil {
		fmt.Printf("read file err=", err)
	}
	fmt.Printf("%v", string(by))
}