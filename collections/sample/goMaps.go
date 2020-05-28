package sample

import "fmt"

func MapSample() {

	//map in go is hashmap
	sampMap := make(map[string]int)
	sampMap2 := map[string]int{"k1": 1, "k2": 2}
	fmt.Println(sampMap2)

	sampMap["k1"] = 1
	sampMap["k2"] = 2
	sampMap["k3"] = 3
	//length
	fmt.Println(len(sampMap))
	//delete
	delete(sampMap, "k1")
	fmt.Println(sampMap)

	//in this 2 value return
	//a indicates the value (default value if the key not present)
	//b boolean value indicate key presence
	a, b := sampMap["k3"]

	fmt.Println(a, " - ", b)

	fmt.Println(sampMap)
	fmt.Println(sampMap["k2"])

}
