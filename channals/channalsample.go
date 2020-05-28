package main

import (
	"fmt"
	"time"
)

func main() {

	src := make(chan int)
	//dst := make(chan int)

	go produce(src)

	go consume(src)

	time.Sleep(10 * time.Second)
}

func consume(dst <-chan int) {
	fmt.Println("from consume")
	for {
		select {
		case val := <-dst:
			fmt.Println("received - ", val)
		default:
			fmt.Println("from default")
		}
	}

}

func produce(src chan int) {
	for _, val := range []int{1, 2, 3, 4, 5, 6, 7} {
		fmt.Println("procuding ")
		src <- val
		time.Sleep(1 * time.Second)
	}
}
