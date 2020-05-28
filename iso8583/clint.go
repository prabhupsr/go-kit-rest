package main

import (
	"fmt"
	"net"
	"os"
	"unsafe"

	"github.com/ideazxy/iso8583"
)

type Data struct {
	bm7  *iso8583.Numeric      `field:"7" length:"10" encode:"bcd"` // bcd value encoding
	bm11 *iso8583.Numeric      `field:"11" length:"6" encode:"bcd"` // bcd value encoding
	bm39 *iso8583.Alphanumeric `field:"39" length:"2"`
	bm70 *iso8583.Numeric      `field:"70" length:"3" encode:"bcd"` // bcd value encoding
}

func main() {
	data := &Data{
		bm11: iso8583.NewNumeric("000145"),
		bm39: iso8583.NewAlphanumeric("00"),
		bm70: iso8583.NewNumeric("019"),
	}
	msg := iso8583.NewMessage("0800", data)
	msg.MtiEncode = iso8583.BCD
	b, err := msg.Bytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("% x\n", b)

	var totallength = len(b) + 2

	lengthBytes := IntToByteArray(totallength)

	var finalBytes []byte

	for _, v := range lengthBytes {
		bytes := append(finalBytes, v)
		fmt.Println(bytes)
	}
	for _, v := range b {
		bytes := append(finalBytes, v)
		fmt.Println(bytes)
	}

	fmt.Println(lengthBytes)
	fmt.Println(b)
	fmt.Println(finalBytes)

	servAddr := "localhost:21886"

	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write(b)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

}

func IntToByteArray(num int) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}
