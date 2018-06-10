package main

import (
	"fmt"
	"log"
	"n2x"
)

func main() {
	n, err := n2x.New("10.71.20.231", "9001")
	if err != nil {
		panic(err)
	}

	err = n.KillSessionByName("SYSTEM")
	if err != nil {
		panic(err)
	}

	sess, err := n.GetSessionByName(n2x.DEFAULTSESSIONNAME)
	if err != nil {
		sess, err = n.OpenNewSession("101/1", "101/2")
		if err != nil {
			panic(err)
		}
	}

	for _, port := range sess.Ports {
		ips, _ := port.LegacyLinkGetSutIPAddresses()
		fmt.Printf("%q ", ips)
	}

	err = sess.StopTest()
	if err != nil {
		panic(err)
	}
	ports, err := sess.GetReservedPorts()
	fmt.Printf("%q ", ports)
	for i, port := range ports {
		port.GetAvailableMediaTypes()
		port.GetMediaType()
		port.SetMediaType(n2x.RJ45)
		port.GetMediaType()
		port.LaserOn()
		port.LaserOff()
		port.LegacyLinkGetSutMacAddress()
		port.LegacyLinkSetSutMacAddress("1.1.1.1", "00:01:00:01:00:02")
		port.LegacyLinkGetSutMacAddress()
		port.LegacyLinkAddSutIPAddress("123.1.1.1")
		//port.LegacyLinkAddSutIPAddresses("120.1.1.5", "30", "1000", "1")
		port.SendAllArpRequests()
		port.SendAllNeighborSolicitations()
		port.LegacyLinkGetAllAddressPools()
		tip := fmt.Sprintf("12.%d.1.1", i)
		sip := fmt.Sprintf("12.%d.1.2", i)
		smac := fmt.Sprintf("00:1%x:00:00:1%x:00", i%255, i%255)
		port.LegacyLinkAddSutIPAddress(sip)
		err := port.LegacyLinkSet("10", tip, "30", smac, sip)
		if err != nil {
			panic(err)
		}
		tip6 := fmt.Sprintf("200%d:1%d::11", i, i)
		sip6 := fmt.Sprintf("200%d:1%d::12", i, i)

		err = port.LegacyLinkSet6("10", tip6, "64", smac, sip6)
		if err != nil {
			panic(err)
		}

		pools, err := port.LegacyLinkGetAllIP6AddressPools()
		if err != nil {
			panic(err)
		}

		port.GetAllOSPFs()
		err = port.DeleteAllOSPFs()
		if err != nil {
			panic(err)
		}
		port.GetAllOSPFs()
		name := fmt.Sprintf("OSPF_EMULATION_%d", i)
		rid := fmt.Sprintf("155.1.1.%d", i)
		srid := fmt.Sprintf("154.1.1.%d", i)
		_, err = port.AddOSPF(fmt.Sprintf("1%d.1%d.1%d.1%d", i, i, i, i), rid, srid, name)
		if err != nil {
			panic(err)
		}
		port.GetAllOSPFs()

		fmt.Printf("%q\n", pools)
	}
	sess.ListModules()
	sess.ListAvailableModules()
	sess.ListAvailablePorts()
	sess.ListLockedPorts()
	sess.ListPorts()
	sess.StopRoutingEngine()
	sess.StartRoutingEngine()
	err = sess.StartTest()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", sess)
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
