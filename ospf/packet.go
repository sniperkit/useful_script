package ospf

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Packet struct {
	EthernetHdr layers.Ethernet
	IPv4Hdr     layers.IPv4
	Header      Header
	Payload     gopacket.Payload
}

type Payload interface {
	IsOSPF()
	Marshal()
	UnMarshal()
	String()
}

func (p *Packet) Marshal() ([]byte, error) {

}

func UnMarshalOSPFPacketFromByte(data []byte) (*Packet, error) {
	p := &Packet{}
}
