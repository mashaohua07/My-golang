package main

import (
	"fmt"
	"time"
)

func main()  {
	//获取当前时间
	now := time.Now()
	//fmt.Printf("now=%v, now=%T\n", now, now)

	//通过now获取年月日，时分秒
	//fmt.Printf("年=%v\n", now.Year())
	//fmt.Printf("月=%v\n", int(now.Month()))
	//fmt.Printf("日=%v\n", now.Day())
	//fmt.Printf("时=%v\n", now.Hour())
	//fmt.Printf("分=%v\n", now.Minute())
	//fmt.Printf("秒=%v\n", now.Second())

	//格式化日期和时间（1）
	fmt.Printf("当前年月日时分秒 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//格式化日期和时间（2）
	fmt.Printf(now.Format("2006-01-02 15:04:05\n"))	//时间日期必须固定2006-01-02 15:04:05
	fmt.Printf(now.Format("01\n"))



	//每隔0.1秒打印一个数，打印到100就退出
	//i := 0
	//for  {
	//	i++
	//	fmt.Println(i)
	//	time.Sleep(time.Millisecond * 100)
	//	if i == 100 {
	//		break
	//	}
	//}


	//Unix和Unixnano的使用
	fmt.Printf("Unix时间戳=%v, Unixnano时间戳=%v\n", now.Unix(), now.UnixNano())
}