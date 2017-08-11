package ospf

import (
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"time"
)

const (
	NormalArea = 1
	STUBArea   = 2
	NSSArea    = 3
)

var allSPFRouters = net.IPAddr{IP: net.IPv4(224, 0, 0, 5)}

type OSPF struct {
	RouterID  net.IP
	AreaID    net.IP
	AreaType  int
	Options   int
	IfName    string
	Conn      *ipv4.RawConn
	Interface *Interface
	CM        *ipv4.ControlMessage
}

func New(ifname string, routerid string, area string, areaType int) (*OSPF, error) {
	instance := &OSPF{
		RouterID: net.ParseIP(routerid),
		AreaID:   net.ParseIP(area),
		AreaType: areaType,
		IfName:   ifname,
	}

	c, err := net.ListenPacket("ip4:89", "0.0.0.0") // OSPF for IPv4
	if err != nil {
		log.Fatal(err)
	}

	r, err := ipv4.NewRawConn(c)
	if err != nil {
		log.Fatal(err)
	}
	instance.Conn = r

	ifp, err := net.InterfaceByName(ifname)
	if err != nil {
		log.Fatal(err)
	}

	oif, err := NewOSPFInterface(ifp)
	if err != nil {

		log.Fatal(err)
	}

	instance.Interface = oif
	instance.CM = &ipv4.ControlMessage{IfIndex: ifp.Index}

	return instance, nil
}

func (o *OSPF) Start() error {
	err := o.Conn.JoinGroup(nil, &allSPFRouters)
	if err != nil {
		log.Printf("Cannot join group: %s\n", err.Error())
		return err
	}
	go o.Hello()
	go o.PacketHandler()
	return nil
}

func (o *OSPF) Stop() error {
	o.Conn.LeaveGroup(nil, &allSPFRouters)
	return nil
}

func (o *OSPF) PacketHandler() {
	buf := make([]byte, 1024)
	for {
		iph, op, cm, err := o.Conn.ReadFrom(buf)
		if err != nil {
			log.Printf("%s", err.Error())
		}

		oh, p, err := UnMarshalOSPFHeader(op)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("oh.Type: ", oh.Type, oh.PacketLength, HeaderLen, len(p))
		c, err := UnMarshalOSPFPacket(p, oh.Type, oh.PacketLength-HeaderLen)
		if err != nil {
			log.Fatal(err)
		}

		switch c.(type) {
		case *Hello:
			h, _ := c.(*Hello)
			h.Header = *oh
			o.ProcessHello(iph, h)
		case *DBD:
			d, _ := c.(*DBD)
			d.Header = *oh
			o.ProcessDBD(iph, d)
		case *LSR:
			lsr, _ := c.(*LSR)
			lsr.Header = *oh
			o.ProcessLSR(iph, lsr)
		case *LSU:
			lsu, _ := c.(*LSU)
			lsu.Header = *oh
			o.ProcessLSU(iph, lsu)
		case *LSAck:
			lsack, _ := c.(*LSAck)
			lsack.Header = *oh
			o.ProcessLSAck(iph, lsack)
		default:
			panic("Unkown ospf packe")
		}
		log.Printf("Received a packet %#v from:\n %s %s %s\n", cm, PrettyIPv4Header(iph), oh, c)
	}
}

func (o *OSPF) Hello() {
	//for time.Duration(time.Second * o.Interface.HelloInterval) {
	for range time.Tick(time.Duration(time.Second * 2)) {
		o.SendHello()
	}
}

func (o *OSPF) SendHello() {
	hello := Packet{
		Header: Header{
			Version:        2,
			Type:           1,
			PacketLength:   24,
			RouterID:       o.RouterID,
			AreaID:         o.AreaID,
			CheckSum:       0,
			AuType:         0,
			Authentication: 0,
		},
		Payload: Hello{
			NetworkMask:            o.Interface.NetworkMask,
			HelloInterval:          o.Interface.HelloInterval,
			Options:                o.Interface.Options,
			RtrPri:                 o.Interface.RouterPriority,
			RouterDeadInterval:     o.Interface.RouterDeadInterval,
			DesignatedRouter:       o.Interface.DesignatedRouter,
			BackupDesignatedRouter: o.Interface.BackupDesignatedRouter,
			Neighbors:              o.Interface.GetAttachedNeighbors(),
		},
	}

	p, err := hello.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	iph := &ipv4.Header{
		Version:  ipv4.Version,
		Len:      ipv4.HeaderLen,
		TOS:      0xc0, // DSCP CS6
		TotalLen: ipv4.HeaderLen + len(p),
		TTL:      1,
		Protocol: 89,
		Dst:      allSPFRouters.IP.To4(),
	}

	if err := o.Conn.WriteTo(iph, p, o.CM); err != nil {
		log.Fatal(err)
	}
}

func (o *OSPF) SendDBD(n *Neighbor) error {
	hello := Packet{
		Header: Header{
			Version:        2,
			Type:           2,
			PacketLength:   24,
			RouterID:       o.RouterID,
			AreaID:         o.AreaID,
			CheckSum:       0,
			AuType:         0,
			Authentication: 0,
		},
		Payload: DBD{
			InterfaceMTU:     1500,
			Options:          0x2,
			MasterSlave:      0x0,
			DDSequenceNumber: n.LastReceivedDBD.DDSequenceNumber,
			LSAHeader:        nil,
		},
	}

	p, err := hello.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	iph := &ipv4.Header{
		Version:  ipv4.Version,
		Len:      ipv4.HeaderLen,
		TOS:      0xc0, // DSCP CS6
		TotalLen: ipv4.HeaderLen + len(p),
		TTL:      1,
		Protocol: 89,
		Dst:      net.IPv4(10, 71, 1, 123),
	}

	if err := o.Conn.WriteTo(iph, p, o.CM); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (o *OSPF) SendLSR() error {
	return nil
}

func (o *OSPF) SendLSU() error {
	return nil
}

func (o *OSPF) SendLSAck() error {
	return nil
}

func (o *OSPF) ProcessHello(iph *ipv4.Header, h *Hello) error {
	n := &Neighbor{}
	n.RouterID = h.Header.RouterID
	n.AreaID = h.Header.AreaID
	n.IP = iph.Src
	n.Options = h.Options
	n.Priority = h.RtrPri
	n.DesignatedRouter = h.DesignatedRouter
	n.BackupDesignatedRouter = h.BackupDesignatedRouter

	ifp := o.Interface

	if _, ok := ifp.Neighbors[n.RouterID.String()]; !ok {
		ifp.Neighbors[n.RouterID.String()] = n
	}

	return nil
}

func (o *OSPF) ProcessDBD(iph *ipv4.Header, d *DBD) error {
	n, ok := o.Interface.Neighbors[d.Header.RouterID.String()]
	if !ok {
		log.Fatal("Cannot find neighbor: %s\n", d.Header.RouterID)
	}
	n.LastReceivedDBD = *d
	//We work as slave, and just confirm to master, if no more lsa from master, stop.
	go o.SendDBD(n)
	return nil
}

func (o *OSPF) ProcessLSR(iph *ipv4.Header, l *LSR) error {
	return nil
}

func (o *OSPF) ProcessLSU(iph *ipv4.Header, l *LSU) error {
	return nil
}

func (o *OSPF) ProcessLSAck(iph *ipv4.Header, l *LSAck) error {
	return nil
}
