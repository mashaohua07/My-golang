package main

import (
	"fmt"
)

func main()  {
	names := []string{"kubernetes", "docker", "istio", "heml", "harbor"}
	heroNames := ""
	fmt.Println("Please enter a keyword")
	fmt.Scanln(&heroNames)

	//for i := 0; i < len(names); i++ {
	//	if heroNames == names[i] {
	//		fmt.Printf("%v\n", i)
	//		break
	//	} else if i == len(names) - 1 {
	//		fmt.Println("Did not match")
	//	}
	//}

	index := -1
	for i := 0; i < len(names); i ++ {
		if heroNames == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("匹配到%v, 下标为%v\n",heroNames, index)
	} else {
		fmt.Println("Did not match")
	}
}
