package main

import (
	"context"
	"fmt"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx"
	"github.com/rsocket/rsocket-go/rx/flux"
	"log"
)

func main() {

	client, err := rsocket.Connect().Resume().Fragment(1024).
		SetupPayload(payload.New([]byte("Hello"), []byte("Hello"))).
		Transport("tcp://127.0.0.1:1221").Start(context.Background())

	if err != nil {
		panic(err)
	}

	//defer client.Close()
	ch := make(chan int, 15)
	chp := make(chan payload.Payload)

	go producer(ch)

	go produce(chp)

	//client.RequestChannel(observable)
	for i := 0; i < 5; i++ {

		fmt.Println(<-ch)
	}
	//go consume(ch)
	//var m rx.Publisher=mono.Just(payload.NewString("abc","asad"))

	errors := make(chan error)
	//var mxx flux.Flux= flux.CreateFromChannel(chp, errors)
	//var m rx.Publisher= flux.CreateFromChannel(chp, errors)

	/*mxx.DoOnNext(func(input payload.Payload) {
		fmt.Println(string(input.Data()))
	}).BlockFirst(context.Background())*/

	var mm rx.Publisher = flux.Create(func(ctx context.Context, s flux.Sink) {

		for i := 0; i < 10; i++ {
			fmt.Println("p ->")
			s.Next(<-chp)
		}
		fmt.Println("complete")
		s.Complete()
	})
	fmt.Print("before sub")
	client.RequestChannel(mm).DoOnNext(func(input payload.Payload) {
		fmt.Println(string(input.Data()))
	}).Subscribe(context.Background())

	result, err := client.RequestResponse(payload.NewString("abcc", "meta")).Block(context.Background())

	if err != nil {
		panic(err)
	}
	log.Println("resp", result)

	fmt.Println("hiiijasd")
	go consume(ch)

	<-errors
}

func produce(chp chan payload.Payload) {
	for i := 0; i < 150; i++ {
		//fmt.Println("producing")
		chp <- payload.NewString("data", "meta")
	}
}

func consume(ch chan int) {
	for i := 0; i < 200; i++ {

		//	fmt.Println(<-ch)
	}
}

func producer(ch chan int) {
	for i := 0; i < 150; i++ {
		ch <- i
	}
}
