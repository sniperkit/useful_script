package ospf

import (
	"errors"
	"time"
)

var errorHeaderTooShort = errors.New("OSPF packet header is to short")
var errorBufferTooShort = errors.New("Packet buffer is too short")

var OSPFBPF string = "proto ospf"

const (
	OSPFVersion2          = 2  // OSPF Version2
	OSPFVersion2HeaderLen = 24 // header length without payload
)

var allSPFRouters = net.IPAddr{IP: net.IPv4(224, 0, 0, 5)}
var allDOtherRouters = net.IPAddr{IP: net.IPv4(224, 0, 0, 6)}

const (
	AreaNormal = 1
	AreaStub   = 2
	AreaNSSA   = 3
)

var (
	Snapshot    int32         = 65535
	Promiscuous bool          = true
	Timeout     time.Duration = 30 * time.Second
)
