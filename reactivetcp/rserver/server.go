package main

import (
	"context"
	"fmt"
	"github.com/rsocket/rsocket-go/rx"
	"github.com/rsocket/rsocket-go/rx/flux"

	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx/mono"
)

func main() {

	err := rsocket.Receive().
		Resume().
		Fragment(1024).
		Acceptor(func(setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
			fmt.Println(string(setup.Data()))

			return rsocket.NewAbstractSocket(reqResp(), reqChannal()), nil
		}).
		Transport("tcp://127.0.0.1:4545").
		Serve(context.Background())

	panic(err)

}

func reqChannal() rsocket.OptAbstractSocket {
	return rsocket.RequestChannel(func(msgs rx.Publisher) flux.Flux {
		fmt.Print("abc")
		flux.Clone(msgs).DoOnNext(func(input payload.Payload) {
			fmt.Print(string(input.Data()))
		}).Subscribe(context.Background())
		return flux.Just(payload.NewString("respp", "cc"))

	})
}

func reqResp() rsocket.OptAbstractSocket {
	return rsocket.RequestResponse(
		func(msg payload.Payload) mono.Mono {
			fmt.Println(string(msg.Data()))
			return mono.Just(msg)
		})
}
