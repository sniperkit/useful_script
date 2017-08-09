package ospf

import (
	"encoding/binary"
	"errors"
	"net"
	"syscall"
)

const (
	Version   = 2  // OSPF Version2
	HeaderLen = 24 // header length without payload
)

var errorHeaderTooShort = errors.New("OSPF packet header is to short")

type Header struct {
	Version        byte
	Type           byte
	PacketLength   int
	RouterID       net.IP
	AreaID         net.IP
	CheckSum       uint16
	AuType         uint16
	Authentication uint64
}

func (h *Header) Marshal() ([]byte, error) {
	if h == nil {
		return nil, syscall.EINVAL
	}
	if h.PacketLength < HeaderLen {
		return nil, errorHeaderTooShort
	}

	b := make([]byte, HeaderLen)
	b[0] = h.Version
	b[1] = h.Type
	binary.BigEndian.PutUint16(b[2:4], uint16(h.PacketLength))
	if ip := h.RouterID.To4(); ip != nil {
		copy(b[4:8], ip[:net.IPv4len])
	}
	if ip := h.AreaID.To4(); ip != nil {
		copy(b[8:12], ip[:net.IPv4len])
	}
	binary.BigEndian.PutUint16(b[12:14], uint16(h.CheckSum))
	binary.BigEndian.PutUint16(b[14:16], uint16(h.AuType))
	binary.BigEndian.PutUint64(b[16:24], uint64(h.Authentication))

	return b, nil
}

func (h *Header) UnMarshal() ([]byte, error) {
	return nil, nil
}

func (h *Header) String() string {
	return ""
}
