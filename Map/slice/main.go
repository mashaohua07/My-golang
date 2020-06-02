package main

import "fmt"

func main()  {
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["new"] = "ren"
		monsters[0]["old"] = "ton"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["new"] = "com"
		monsters[1]["old"] = "cn"
	}

	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string)
	//	monsters[2]["string"] = "str"
	//	monsters[2]["itin"] = "it"
	//}

	newMosters := make(map[string]string)
	newMosters["string"] = "str"
	newMosters["init"] = "int"
	monsters = append(monsters, newMosters)
	fmt.Println(monsters)

}