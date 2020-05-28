package main

import (
	"errors"
	"fmt"
	_ "time"
)

func test()  {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover()	//内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()
	i := 10
	j := 0
	res := i / j
	fmt.Println("res = ", res)
}


//函数去读取一个ini.conf的配置文件
//如果文件名传入不正确，我们就返回一个自定义错误
func readConf(name string) (err error)  {
	if name == "ini.conf" {
		return nil
	}else {
		return errors.New("读取文件名错误")
	}
}

func test2()  {
	err := readConf("ini.conf")
	if err != nil {
			panic(err)
	}
	fmt.Println("test2()下面的代码...")
}

func main()  {
	test()

	//测试自定义错误的使用
	test2()

	//for  {
		fmt.Println("main()下面的代码...")
	//	time.Sleep(time.Second)
	//}

}
