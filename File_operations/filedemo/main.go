package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
)

func main()  {
	file, err := os.Open("../test")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for  {
		str, err := reader.ReadString('\n')
		if err == io.EOF {	//io.EOF代表文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")
}