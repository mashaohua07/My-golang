package main

import "fmt"

func main()  {
	var a map[string]string
	a = make(map[string]string)
	a["node"] = "pod"
	a["var"] = "pdd"
	a["sins"] = "old"
	fmt.Println(a)

	b := make(map[string]string)
	b["string"] = "word"
	b["latest"] = "fint"
	fmt.Println(b)

	c := map[string]string{
		"int" : "in",
		"ina" : "and",
	}
	fmt.Println(c)
	fmt.Println()

	d := make(map[string]map[string]string)
	d["v"] = make(map[string]string, 2)
	d["v"]["t"] = "s"
	d["v"]["d"] = "v"
	d["b"] = make(map[string]string, 2)
	d["b"]["d"] = "v"
	d["b"]["t"] = "s"
	fmt.Println(d)
	fmt.Println(d["v"])
	fmt.Println(d["v"]["t"])
}