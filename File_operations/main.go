package main

import (
	"os"
	"fmt"
)

func main()  {
	file , err := os.Open("test")
	if err != nil {
		fmt.Println("open dile err=", err)
	}

	fmt.Printf("file=%v\n", file)

	err = file.Close()
	if err != nil {
		fmt.Println("clone file err=", err)
	}
}