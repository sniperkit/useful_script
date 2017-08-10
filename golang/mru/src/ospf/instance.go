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

var (
	device       string = "ens33"
	snapshot_len int32  = 65535
	promiscuous  bool   = true
	err          error
	timeout      time.Duration = 40 * time.Second
	rxHandler    *pcap.Handle
	txHandler    *pcap.Handle
)

func (o *OSPF) Start() error {
	txHandler, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if sendHdl == nil {
		server.logger.Err(fmt.Sprintln("SendHdl: No device found.", ent.IfName, err))
		return
	}

	rxHandler, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}

	//go o.Hello()
	go o.PacketHandler()

	return nil
}

func (o *OSPF) Stop() error {
	o.Conn.LeaveGroup(o.Interface.Interface, &allSPFRouters)
	return nil
}

func (o *OSPF) PacketHandler() {
	// Set filter
	//var filter string = "tcp and port 22"
	var filter string = "proto ospf and not src host 10.71.1.122" //Try to capture ospf packet

	err = rxHandler.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	packetSource := gopacket.NewPacketSource(rxHandler, layers.LayerTypeEthernet)
	for pkt := range packetSource.Packets() {
		// Do something with a packet here.
		fmt.Println(pkt)
		ethLayer := pkt.Layer(layers.LayerTypeEthernet)
		if ethLayer == nil {
			log.Println("Not an Ethernet frame")
			return
		}
		eth := ethLayer.(*layers.Ethernet)

		ipLayer := pkt.Layer(layers.LayerTypeIPv4)
		if ipLayer == nil {
			log.Println("Not an IP packet")
			return
		}

		ospfPkt := ipLayer.LayerPayload()
		ospfData := ospfPkt[24:]
		log.Printf("eth: %#v\n", eth)
		log.Printf("ip:  %#v\n", ipLayer)
		log.Printf("ospf: %#v\n", ospfPkt)
		log.Printf("ospfdata: %#v\n", ospfData)
	}

}

func (o *OSPF) Hello() {
	//for time.Duration(time.Second * o.Interface.HelloInterval) {
	for range time.Tick(time.Duration(time.Second * 10)) {
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
			RtrPri:                 o.Interface.Priority,
			RouterDeadInterval:     o.Interface.DeadInterval,
			DesignatedRouter:       o.Interface.DesignatedRouter,
			BackupDesignatedRouter: o.Interface.BackupDesignatedRouter,
			Neighbors:              o.Interface.GetAllAttachedNeighbors(),
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

	err = txHandler.WritePacketData()
	if err := o.Conn.WriteTo(iph, p, o.CM); err != nil {
		log.Fatal(err)
	}
}

func BuildHelloPkt(ifp *net.Interface) []byte {
	ospfHdr := OSPFHeader{
		ver:      OSPF_VERSION_2,
		pktType:  uint8(HelloType),
		pktlen:   0,
		routerId: server.ospfGlobalConf.RouterId,
		areaId:   ent.IfAreaId,
		chksum:   0,
		authType: ent.IfAuthType,
		//authKey:        ent.IfAuthKey,
	}

	/*
	   Rfc 2328 4.5
	   ExternalRoutingCapability
	           Entire OSPF areas can be configured as "stubs" (see Section
	           3.6).  AS-external-LSAs will not be flooded into stub areas.
	           This capability is represented by the E-bit in the OSPF
	           Options field (see Section A.2).  In order to ensure
	           consistent configuration of stub areas, all routers
	           interfacing to such an area must have the E-bit clear in
	           their Hello packets
	*/
	if ent.IfAreaId == nil {
		log.Printf("HELLO: Null area id for intfkey %v", ent)
		return nil
	}
	areaId := config.AreaId(convertIPInByteToString(ent.IfAreaId))
	isStub := server.isStubArea(areaId)
	option := uint8(2)
	if isStub {
		option = uint8(0)
	}
	helloData := OSPFHelloData{
		netmask:             ent.IfNetmask,
		helloInterval:       ent.IfHelloInterval,
		options:             option,
		rtrPrio:             ent.IfRtrPriority,
		rtrDeadInterval:     ent.IfRtrDeadInterval,
		designatedRtr:       ent.IfDRIp,
		backupDesignatedRtr: ent.IfBDRIp,
		//designatedRtr:          []byte {0, 0, 0, 0},
		//backupDesignatedRtr:    []byte {0, 0, 0, 0},
		//neighbor:               []byte {1, 1, 1, 1},
	}

	var neighbor []byte
	var nbrlen = 0
	nbr := make([]byte, 4)
	for key, _ := range ent.NeighborMap {
		nbrConf := server.NeighborConfigMap[key]
		binary.BigEndian.PutUint32(nbr, nbrConf.OspfNbrRtrId)
		nbrlen = nbrlen + 4
		neighbor = append(neighbor, nbr...)
	}

	ospfPktlen := OSPF_HEADER_SIZE
	ospfPktlen = ospfPktlen + OSPF_HELLO_MIN_SIZE + nbrlen

	ospfHdr.pktlen = uint16(ospfPktlen)

	ospfEncHdr := encodeOspfHdr(ospfHdr)
	//server.logger.Debug(fmt.Sprintln("ospfEncHdr:", ospfEncHdr))
	helloDataEnc := encodeOspfHelloData(helloData)
	//server.logger.Debug(fmt.Sprintln("HelloPkt:", helloDataEnc))
	helloDataNbrEnc := append(helloDataEnc, neighbor...)
	//server.logger.Debug(fmt.Sprintln("HelloPkt with Neighbor:", helloDataNbrEnc))

	ospf := append(ospfEncHdr, helloDataNbrEnc...)
	//server.logger.Debug(fmt.Sprintln("ospf:", ospf))
	csum := computeCheckSum(ospf)
	binary.BigEndian.PutUint16(ospf[12:14], csum)
	copy(ospf[16:24], ent.IfAuthKey)

	ipPktlen := IP_HEADER_MIN_LEN + ospfHdr.pktlen
	ipLayer := layers.IPv4{
		Version:  uint8(4),
		IHL:      uint8(IP_HEADER_MIN_LEN),
		TOS:      uint8(0xc0),
		Length:   uint16(ipPktlen),
		TTL:      uint8(1),
		Protocol: layers.IPProtocol(OSPF_PROTO_ID),
		SrcIP:    ent.IfIpAddr,
		DstIP:    net.IP{224, 0, 0, 5},
	}

	ethLayer := layers.Ethernet{
		SrcMAC:       ent.IfMacAddr,
		DstMAC:       net.HardwareAddr{0x01, 0x00, 0x5e, 0x00, 0x00, 0x05},
		EthernetType: layers.EthernetTypeIPv4,
	}

	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	gopacket.SerializeLayers(buffer, options, &ethLayer, &ipLayer, gopacket.Payload(ospf))
	//server.logger.Debug(fmt.Sprintln("buffer: ", buffer))
	ospfPkt := buffer.Bytes()
	//server.logger.Debug(fmt.Sprintln("ospfPkt: ", ospfPkt))
	return ospfPkt
}
