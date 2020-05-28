package main

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/utils"
	"os"
	"strings"
	"time"

	"github.com/go-netty/go-netty/codec/frame"
	"github.com/go-netty/go-netty/transport/tcp"
)

func main() {

	fmt.Println("abc-->>>", []byte("abc"))

	// child pipeline initializer
	var childPipelineInitializer = func(channel netty.Channel) {

		fmt.Println("childPipelineInitializer")
		channel.Pipeline().
			// the maximum allowable packet length is 128 bytes，use \n to splite, strip delimiter
			AddLast(frame.DelimiterCodec(128, "\n", true)).
			// convert to string
			AddLast(InboundAdapter{}).
			// LoggerHandler, print connected/disconnected event and received messages
			AddLast(LoggerHandler{}).
			// UpperHandler (string to upper case)
			AddLast(UpperHandler{})
	}

	// new go-netty bootstrap
	/*netty.NewBootstrap().
	// configure the child pipeline initializer
	ChildInitializer(childPipelineInitializer).
	// configure the transport protocol
	Transport(tcp.New()).
	// configure the listening address
	Listen("0.0.0.0:9527").
	// waiting for exit signal
	Action(netty.WaitSignal(os.Interrupt)).
	// print exit message
	Action(func(bs netty.Bootstrap) {
		fmt.Println("server exited")
	})
	*/
	tcpOptions := &tcp.Options{
		Timeout:         time.Second * 3,
		KeepAlive:       true,
		KeepAlivePeriod: time.Second * 5,
		Linger:          0,
		NoDelay:         true,
		SockBuf:         1024,
	}

	bootstrap := netty.NewBootstrap()
	bootstrap.Channel(netty.NewBufferedChannel(128, 1024))
	bootstrap.ChildInitializer(childPipelineInitializer).ClientInitializer(childPipelineInitializer)
	bootstrap.ChannelExecutor(netty.NewFixedChannelExecutor(128, 1))
	bootstrap.Transport(tcp.New())
	bootstrap.Listen("127.0.0.1:9527", tcp.WithOptions(tcpOptions))

	c, err := bootstrap.Connect("tcp://localhost:9527", "go-netty")
	if nil != err {
		panic(err)
	}

	i, err := c.Write([]byte("msgghgchghgcghvjhvjfhjgfh gfsgffgh "))

	if err != nil {
		panic(err)
	}

	fmt.Println(i)

	netty.NewBootstrap().
		// configure the child pipeline initializer
		ChildInitializer(childPipelineInitializer).
		// configure the transport protocol
		Transport(tcp.New()).
		// configure the listening address
		Listen("0.0.0.0:9927").
		// waiting for exit signal
		Action(netty.WaitSignal(os.Interrupt)).
		// print exit message
		Action(func(bs netty.Bootstrap) {
			fmt.Println("server exited")
		})

	/*netty.NewBootstrap().
	// configure the child pipeline initializer
	ChildInitializer(childPipelineInitializer).
	// configure the transport protocol
	Transport(tcp.New()).
	// configure the listening address
	Listen("0.0.0.0:9527").
	ChannelExecutor(NewFixedChannelExecutor(128, 1))
	// waiting for exit signal
	Action(netty.WaitSignal(os.Interrupt)).
	// print exit message
	Action(func(bs netty.Bootstrap) {
		fmt.Println("server exited")
	})*/

}

type LoggerHandler struct{}

func (LoggerHandler) HandleActive(ctx netty.ActiveContext) {
	fmt.Println("go-netty:", "->", "active:", ctx.Channel().RemoteAddr())
	// write welcome message
	ctx.Write("Hello I'm " + "go-netty")
}

func (LoggerHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	fmt.Println("go-netty:", "->", "handle read:", message)
	// leave it to the next handler(UpperHandler)
	ctx.HandleRead(message)
}

func (LoggerHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	fmt.Println("go-netty:", "->", "inactive:", ctx.Channel().RemoteAddr(), ex)
	// disconnected，the default processing is to close the connection
	ctx.HandleInactive(ex)
}

type UpperHandler struct{}

func (UpperHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	// text to upper case
	text := message.(string)
	upText := strings.ToUpper(text)
	// write the result to the client
	ctx.Write(text + " -> " + upText)
}

type InboundAdapter struct{}

func (InboundAdapter) HandleRead(ctx netty.InboundContext, message netty.Message) {
	// text to upper case
	textBytes := utils.MustToBytes(message)

	// convert from []byte to string
	sb := strings.Builder{}
	sb.Write(textBytes)

	// post text
	ctx.HandleRead(sb.String())
}
