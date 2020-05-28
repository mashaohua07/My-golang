package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	//golang的编码统一为utf-8,(ascii的字符占一个字节，汉子占3个字节)
	str := "hello上海"
	fmt.Println("长度:", len(str))

	str2 := "hello上海"
	//字符串遍历，同时处理有中文的问题 r := []rune(str)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符: %c\n", r[i])
	}

	//字符串转整数： n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("456")
	if err != nil {
		fmt.Println("转换格式错误", err)
	}else {
		fmt.Println("转换成的格式是:", n)
	}

	//整数转字符串： str = strconv.Itoa("12345")
	str = strconv.Itoa(12345)
	fmt.Printf("str=%T, str=%v\n", str, str)


	//字符串转[]byte
	var bytes = []byte("hello")
	fmt.Printf("bytes=%v\n", bytes)


	//[]byte转字符串
	str = string([]byte{104, 101, 108})
	fmt.Printf("str=%v\n", str)


	//10进制转 2， 8， 16进制
	str = strconv.FormatInt(123, 2)
	fmt.Println("123对应的二进制是:", str)
	str = strconv.FormatInt(123, 16)
	fmt.Println("123对应的十六进制是:", str)


	//查找子串是否在指定的字符串中
	b := strings.Contains("kubernetes", "ber")
	fmt.Printf("b=%v\n", b)


	//统计一个字符串中有几个指定的字串
	num := strings.Count("kubernetes", "e")
	fmt.Printf("num = %v\n", num)

	//不区分大小写的字符串比较(==是区分大小写的)
	c := strings.EqualFold("abc", "abc")
	fmt.Printf("c = %v\n", c)


	//返回子串在字符串第一次出现的index值，如果没有就返回-1
	index := strings.Index("kubernetes", "net")
	fmt.Printf("index = %v\n", index)


	//返回子串在字符串最后一次出现的index值，如果没有就返回-1
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("index = %v\n", index)


	//将指定的字符串替换成另一个子串
	//n可以指定你想替换的个数，-1表示全部替换
	str2 = "go goland golang"
	str = strings.Replace(str2, "go", "go语言", 1)
	fmt.Printf("str2 = %v, str = %v\n", str2, str)


	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split("kubernetes,docker,istio,jenkins", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr = %v\n", strArr)


	//将字符串的字母进行大小写的转换
	str = "Kubernetes Jenkins"
	str = strings.ToLower(str)
	//str = strings.ToUpper(str)  //全部大写
	fmt.Printf("str = %v\n", str)


	//将字符串左右俩边的空格去掉
	str = strings.TrimSpace(" master and node && orgin   ")
	fmt.Printf("str = %q\n", str)


	//将字符串左右俩边的指定字符去掉
	str = strings.Trim("! vender ", " !")
	fmt.Printf("str = %q\n", str)


	//判断字符串是否以指定的字符串开头[strings.HasSuffix是判断结尾]
	b = strings.HasPrefix("ftp://172.17.16.16", "ftp")
	fmt.Printf("b = %v\n", b)
}
