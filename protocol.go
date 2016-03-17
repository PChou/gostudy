package main

import (
	"encoding/binary"
	"errors"
)

var (
	invalidHeadLenError = errors.New("invalid head len")
	invalidPacketError  = errors.New("invalid packet len")
)

type SimplePacketProtocol struct {
}

func (protocol *SimplePacketProtocol) HeadSize() uint8 {
	return 4
}

func (protocol *SimplePacketProtocol) ValidHead(head []byte, n uint8) (bodySize uint32, err error) {
	if n != 4 && len(head) != 4 {
		return 0, invalidHeadLenError
	}

	if head[0] != 77 || head[1] != 67 {
		return 0, invalidHeadLenError
	}

	len := binary.BigEndian.Uint16(head[2:])
	return uint32(len), nil
}

func (protocol *SimplePacketProtocol) UnBoxing(packet []byte, n uint32) (body []byte, err error) {
	if len(packet) > 4 {
		return packet[4:], nil
	}

	return nil, invalidPacketError
}

func (protocol *SimplePacketProtocol) Boxing(body []byte, n uint32) (packet []byte, pn uint32, err error) {
	len := len(body)
	lenShort := uint16(len)
	buf := make([]byte, len+2)
	binary.BigEndian.PutUint16(buf[0:2], lenShort)
	copy(buf[0:], buf)
	copy(buf[2:], body)
	return buf, uint32(len + 2), nil
}

func NewPacketProtocol() PacketProtocol {
	return &SimplePacketProtocol{}
}
