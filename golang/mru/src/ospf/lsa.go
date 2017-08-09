package ospf

import (
	"net"
)

const (
	Router         = 1
	Network        = 2
	NetworkSummary = 3
	ASBRSummary    = 4
	ASExternal     = 5
)

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
