package ospf

import (
	"encoding/binary"
	"errors"
	"net"
	"syscall"
)

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
