package ospf

import (
	"net"
)

type Neighbor struct {
	RouterID net.IP
	AreadID  net.IP
	AreaType int
}
