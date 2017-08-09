package ospf

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
)

type Hello struct {
	NetworkMask            net.IPMask
	HelloInterval          uint16
	Options                uint8
	RtrPri                 uint8
	RouterDeadInterval     uint32
	DesignatedRouter       net.IP
	BackupDesignatedRouter net.IP
	Neighbors              []net.IP
}

func (h Hello) Marshal() ([]byte, error) {
	b := make([]byte, HeaderLen+4*len(h.Neighbors))
	copy(b[0:4], h.NetworkMask)

	binary.BigEndian.PutUint16(b[4:6], uint16(h.HelloInterval))
	b[6] = h.Options
	b[7] = h.RtrPri
	binary.BigEndian.PutUint32(b[8:12], uint32(h.RouterDeadInterval))

	if ip := h.DesignatedRouter.To4(); ip != nil {
		copy(b[12:16], ip[:net.IPv4len])
	}

	if ip := h.BackupDesignatedRouter.To4(); ip != nil {
		copy(b[16:20], ip[:net.IPv4len])
	}

	for i, n := range h.Neighbors {
		if ip := n.To4(); ip != nil {
			copy(b[20+i*4:20+(i+1)*4], ip[:net.IPv4len])
		}
	}

	return b, nil
}

func (h Hello) IsOSPF() bool {
	return true
}

func (h Hello) UnMarshal(data []byte) (interface{}, error) {
	return nil, nil
}

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
