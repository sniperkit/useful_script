package main

import (
	"fmt"
	"log"
	"n2x"
	//	"time"
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
		panic(err)
	}

	for _, port := range sess.Ports {
		ips, _ := port.LegacyLinkGetSutIPAddresses()
		fmt.Printf("%q ", ips)
	}

	ports, err := sess.GetReservedPorts()
	fmt.Printf("%q ", ports)
	sess.StopRoutingEngine()
	//	time.Sleep(time.Duration(time.Second * 40))
	for i, port := range ports {
		name := fmt.Sprintf("OSPF_EMULATION_%d", i)
		//	port.GetAllOSPFs()

		ospf, err := port.GetOSPFByName(name)
		if err != nil {
			panic(err)
		}
		err = ospf.RemoveAllExternalRoutePools()
		if err != nil {
			panic(err)
		}

		err = ospf.RemoveAllPools()
		if err != nil {
			panic(err)
		}

		_, err = ospf.AddRouter()
		if err != nil {
			panic(err)
		}
		_, err = ospf.AddNetwork()
		if err != nil {
			panic(err)
		}

		_, err = ospf.AddSummaryRoutePool()
		if err != nil {
			panic(err)
		}

		ep, err := ospf.AddExternalRoutePool()
		if err != nil {
			panic(err)
		}

		_, err = ep.GetExternalLsaPool()
		if err != nil {
			panic(err)
		}

		err = ep.SetRoutes("123.1.1.0", "32", "10000", "1")
		if err != nil {
			panic(err)
		}

		err = ospf.AddExternalPrefixes("100.1.1.1", "32", "10000", "1", "1", "1", "0.0.0.0")
		if err != nil {
			panic(err)
		}

		err = ospf.AddExternalPrefixes("101.1.1.1", "32", "10000", "1", "1", "1", "0.0.0.0")
		if err != nil {
			panic(err)
		}

		err = ospf.AddExternalPrefixes("102.1.1.1", "32", "10000", "1", "1", "1", "0.0.0.0")
		if err != nil {
			panic(err)
		}

		_, err = ospf.GetObjectConnections()
		if err != nil {
			panic(err)
		}

		err = ospf.WithdrawAllExternalPrefixes()
		if err != nil {
			panic(err)
		}

		/*
			err = ospf.AdvertiseAllExternalPrefixes()
			if err != nil {
				panic(err)
			}
		*/
	}

	sess.StartRoutingEngine()
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
