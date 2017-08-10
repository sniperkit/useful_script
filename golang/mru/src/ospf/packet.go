package ospf

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
)

type Packet struct {
	Header  Header
	Payload Payload
}

func (p *Packet) Marshal() ([]byte, error) {
	hdr, err := p.Header.Marshal()
	if err != nil {
		log.Println("Header Marshal failed\n")
		panic(err)
	}

	body, err := p.Payload.Marshal()
	if err != nil {
		log.Println("Payload Marshal failed\n")
		panic(err)
	}

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, hdr)
	binary.Write(&buffer, binary.BigEndian, body)
	p.Header.PacketLength = buffer.Len()
	buffer.Reset()
	hdr, err = p.Header.Marshal()
	if err != nil {
		log.Println("Header Checksum Marshal failed\n")
		panic(err)
	}
	binary.Write(&buffer, binary.BigEndian, hdr)
	binary.Write(&buffer, binary.BigEndian, body)
	p.Header.CheckSum = CheckSum(buffer.Bytes())

	hdr, err = p.Header.Marshal()
	if err != nil {
		log.Println("Header Checksum Marshal failed\n")
		panic(err)
	}
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, hdr)
	binary.Write(&buffer, binary.BigEndian, body)

	return buffer.Bytes(), nil
}

func (p *Packet) UnMarshal() (interface{}, error) {
	return nil, nil
}

type Payload interface {
	IsOSPF() bool
	Marshal() ([]byte, error)
	UnMarshal([]byte) (interface{}, error)
}
