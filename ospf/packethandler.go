package ospf

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

type PacketHandler struct {
	rx *gopacket.PacketSource
	tx *pcap.Handle
}

func (ph *PacketHandler) Send(p []byte) error {
	err := ph.tx.WritePacketData(p)
	if err != nil {
		return err
	}
}

func (ph *PacketHandler) Recv() <-chan *Packet {
	ch := make(chan *Packet, 10)
	go func(ch chan<- *Pakcet) {
		for p := range ph.rx {
			op, err := ph.ParseOSPFPacket(p)
			if err != nil {
				log.Printf("Cannot parse received packet: %s", err.Error())
				continue
			}
			ch <- op
		}
	}(ch)

	return ch
}

func (ph *PacketHandler) ParseOSPFPacket(gp *gopacket.Packet) (*Packet, error) {
	ethLayer := gp.Layer(layers.LayerTypeEthernet)
	if ethLayer == nil {
		log.Println("Not an Ethernet frame")
		return nil, err
	}
	eth := ethLayer.(*layers.Ethernet)

	ipLayer := gp.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		log.Println("Not an IP packet")
		return nil, err
	}

	iph, err := ParaseIPv4Header(ipLayer.LayerContent())
	if err != nil {
		log.Printf("Cannot parse IP header: %s\n", err.Error())
		return nil, err
	}

	ipLayer.LayerPayload()
	op, err := ParseOSPFPacket(ipLayer.LayerPayload())
	if err != nil {
		log.Printf("Cannot parse OSPF Packet: %s\n", err.Error())
		return nil, err
	}

	p := &Packet{}
	P.IPv4Header = *iph
	p.Header = op.Header
	p.Payload = op.Payload

	return p, nil
}

func NewPacketHandler(ifp *Interface) (*PacketHandler, error) {
	hdr := &PacketHandler{}
	handle, err := pcap.OpenLive(ifp.Name, Snapshot, Promiscuous, Timeout)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return nil, err
	}

	err = handle.SetBPFFilter(OSPFBPF)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return nil, err
	}

	rxSource := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)

	hdr.tx = handle
	hdr.rx = rxSource

	return hdr, nil
}
