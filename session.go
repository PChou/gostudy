package main

import (
//"fmt"
)

// interface between session and protocal
// the implement should indicate how to recv an entire packet
type PacketProtocol interface {
	// indicate head size of an valid packet
	HeadSize() uint8

	//validate the head data
	//return body size
	ValidHead(head []byte, n uint8) (bodySize uint32, err error)

	//validate the packet data, and return body data
	//the packet data is an entire packet include head and body
	UnBoxing(packet []byte, n uint32) (body []byte, err error)

	//boxing the body into a package and used to send
	Boxing(body []byte, n uint32) (packet []byte, pn uint32, err error)
}
