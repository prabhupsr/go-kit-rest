package sample

import "fmt"

func ArraySample() {
	//Array

	var arr []int = []int{1, 2, 3, 4, 5}

	printArray("arr", arr)

	toArr := make([]int, 5)

	copy(toArr, arr)
	printArray("toArr", toArr)

	//array  slice can be created with make with default values
	makeArr := make([]int, 3)
	for i := 0; i < 3; i++ {
		makeArr[i] = i * i
	}

	printArray("makeArr ", makeArr)

	//slicing
	//arrayname[startindex(default 0):endindex (default (len(arr)))]
	printArray("arr[1:] ", arr[1:])
	printArray("arr[:2] ", arr[:2])
	printArray("arr[1:2] ", arr[1:2])

}

func printArray(msg string, arr []int) {
	fmt.Println(msg, "-", arr)
}
