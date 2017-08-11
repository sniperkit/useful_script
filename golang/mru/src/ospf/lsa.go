package ospf

import (
	"encoding/binary"
	"fmt"
	"net"
)

const (
	Router         = 1
	Network        = 2
	NetworkSummary = 3
	ASBRSummary    = 4
	ASExternal     = 5
)

type LSA interface {
}

type LSAIdentity struct {
	LSType            uint32
	LSID              net.IP
	AdvertisingRouter net.IP
}

type LSAHeader struct {
	LSAge             uint16
	Options           uint8
	LSType            uint8
	LinkStateID       net.IP
	AdvertisingRouter net.IP
	LSSequenceNumber  uint32
	LSCheckSum        uint16
	Length            uint16
}

func UnMarshalLSAHeader(b []byte) (*LSAHeader, error) {
	lsh := &LSAHeader{}
	lsh.LSAge = binary.BigEndian.Uint16(b[0:2])
	lsh.Options = b[2]
	lsh.LSType = b[3]
	lsh.LinkStateID = net.IPv4(b[4], b[5], b[6], b[7])
	lsh.AdvertisingRouter = net.IPv4(b[8], b[9], b[10], b[11])
	lsh.LSSequenceNumber = binary.BigEndian.Uint32(b[12:16])
	lsh.LSCheckSum = binary.BigEndian.Uint16(b[16:18])
	lsh.Length = binary.BigEndian.Uint16(b[18:20])

	return lsh, nil
}

var LSATypeToName = map[uint8]string{
	1: "Router LSA",
	2: "Network LSA",
	3: "Network Summary LSA",
	4: "ASBR Summary LSA",
	5: "AS External LSA",
}

func (lsh *LSAHeader) String() string {
	var s string
	s += fmt.Sprintf("----------------------------------------------------------------------\n")
	s += fmt.Sprintf("        Type:                   :%d(%s)\n", lsh.LSType, LSATypeToName[lsh.LSType])
	s += fmt.Sprintf("        LinkStateID             : %s\n", lsh.LinkStateID)
	s += fmt.Sprintf("        AdvertisingRouter       : %s\n", lsh.AdvertisingRouter)
	s += fmt.Sprintf("        Age                     : %d\n", lsh.LSAge)
	s += fmt.Sprintf("        LSCheckSum              : 0x%x\n", lsh.LSCheckSum)
	s += fmt.Sprintf("        LSSequenceNumber        : 0x%x\n", lsh.LSSequenceNumber)
	s += fmt.Sprintf("        Options                 : 0x%x\n", lsh.Options)
	s += fmt.Sprintf("        Length                  : %d\n", lsh.Length)

	return s
}

type LinkState struct {
	LinkID      net.IP
	LinkData    net.IP
	Type        uint8
	NumberOfTOS uint8
	TOS0Metric  uint16
}

type NetworkLSA struct {
	LSAHeader
	NetworkMask    net.IP
	AttachedRouter []net.IP
}

type NetworkSum struct {
	NetworkMask net.IP
	metric      uint32 /* 24 bit by standard */
	TOS         uint8
	TOSMetric   uint32 /* 24 bit by standard */
}

type NetworkSummaryLSA struct {
	LSAHeader
	Networks []NetworkSum
}

type ASBRSum struct {
	NetworkMask net.IP
	metric      uint32 /* 24 bit by standard */
	TOS         uint8
	TOSMetric   uint32 /* 24 bit by standard */
}

type ASBRSummaryLSA struct {
	LSAHeader
	ASBRs []ASBRSum
}

type ExternalNetwork struct {
	NetworkMask       net.IP
	EBit              bool
	Metric            uint32 /* 24 bit by standard */
	ForWardingAddress net.IP
	ExternalRouteTag  uint32
}

type ASExternalLSA struct {
	LSAHeader
	Networks []ExternalNetwork
}
