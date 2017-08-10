package ospf

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
)

type Hello struct {
	Header                 Header
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
