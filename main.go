package main

import (
	"fmt"
	"net"
)

func doConnect(conn net.Conn, protocol PacketProtocol) {
	defer conn.Close()

	hz := protocol.HeadSize()
	buf := make([]byte, hz)
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("read head failed.")
			return
		}

		bs, err := protocol.ValidHead(buf, uint8(n))
		if err != nil {
			fmt.Println("protocol error In HeadValid .")
			return
		}

		fmt.Println("len of body from head is ", bs)
		buf2 := make([]byte, bs)
		n, err = conn.Read(buf2[0:])
		if err != nil {
			fmt.Println("read body failed.", err.Error())
			return
		}

		buf3 := make([]byte, len(buf)+len(buf2))
		copy(buf3[0:], buf)
		copy(buf3[len(buf):], buf2)

		fmt.Println(buf3)

		body, err := protocol.UnBoxing(buf3, uint32(len(buf3)))
		if err != nil {
			fmt.Println("unboxing body failed.")
			return
		}
		fmt.Println(string(body))
	}

}

func main() {
	fmt.Println("Starting server...")

	//test := [4]byte{'M', 'C', 0x01, 0x02}
	//for _, v := range test {
	//	fmt.Println(v)
	//}

	//in := binary.LittleEndian.Uint16(test[2:])
	//fmt.Println(in)

	//if test[0] == 77 {
	//	fmt.Println("aa")
	//}

	//fmt.Println(rune(test[0]))

	//fmt.Println(test[0:])
	//if string(test[0:1]) == "M" {
	//	fmt.Println("MMM")
	//}

	protocol := NewPacketProtocol()
	listener, err := net.Listen("tcp", "127.0.0.1:10004")
	if err != nil {
		fmt.Println("listen 127.0.0.1:10004 failed.")
		return
	}
	defer listener.Close()
	fmt.Println("begin listen at 127.0.0.01:1004")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept client error.")
			return
		}

		fmt.Println("comming a client..")
		go doConnect(conn, protocol)
	}

}
