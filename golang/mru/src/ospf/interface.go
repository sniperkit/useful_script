package ospf

import (
	"errors"
	"net"
)

type Interface struct {
	Name                   string
	Priority               uint8
	HelloInterval          uint16
	DeadInterval           uint32
	RetransmitInterval     int
	TransmitDelay          int
	NetworkType            int
	Options                uint8
	Interface              *net.Interface
	IP                     net.IP
	NetworkMask            net.IPMask
	DesignatedRouter       net.IP
	BackupDesignatedRouter net.IP
	Neighbors              []*Neighbor
}

func NewOSPFInterface(ifp *net.Interface) (*Interface, error) {
	newIfp := Interface{
		Name:               ifp.Name,
		Priority:           1,
		HelloInterval:      10,
		DeadInterval:       40,
		RetransmitInterval: 10,
		TransmitDelay:      1,
		NetworkType:        1,
		Interface:          ifp,
	}

	ip, mask, err := GetIPv4Address(ifp)
	if err != nil {
		panic(err)
	}

	newIfp.IP = *ip
	newIfp.NetworkMask = *mask

	return &newIfp, nil
}

func GetIPv4Address(ifp *net.Interface) (*net.IP, *net.IPMask, error) {
	var ip *net.IP
	var mask *net.IPMask
	addrs, _ := ifp.Addrs()
	for _, a := range addrs {
		r, n, err := net.ParseCIDR(a.String())
		if err != nil {
			continue
		}

		if r.To4() != nil {
			ip = &r
			mask = &n.Mask

			return ip, mask, nil
		}
	}

	return nil, nil, errors.New("Interface has no ipv4 addresss")
}

func (oi *Interface) Start() error {
	return nil
}

func (oi *Interface) Stop() error {
	return nil
}

func (oi *Interface) GetAllAttachedNeighbors() []net.IP {
	ns := make([]net.IP, 0, len(oi.Neighbors))
	for _, n := range oi.Neighbors {
		ns = append(ns, n.RouterID)
	}

	return ns
}
