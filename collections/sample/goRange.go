package sample

import "fmt"

func RangeSample() {

	//range for array
	for index, value := range []int{1, 3, 4, 6, 2} {
		fmt.Println(index, value)
	}

	//range for map
	for key, value := range map[int]string{1: "a", 2: "b"} {
		fmt.Println(key, value)
	}

	//range for string
	for pos, c := range "collection" {
		fmt.Println(pos, c)
	}

}
