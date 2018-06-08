package main

import (
	"fmt"
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
		ips, _ := port.LegacyGetSutIPAddresses()
		fmt.Printf("%q ", ips)
	}

	ports, err := sess.GetReservedPorts()
	fmt.Printf("%q ", ports)
	modules, _ := sess.ListModules()
	fmt.Printf("%q ", modules)
	modules, _ = sess.ListAvailableModules()
	fmt.Printf("%q ", modules)
	modules, _ = sess.ListAvailablePorts()
	fmt.Printf("%q ", modules)
	modules, _ = sess.ListLockedPorts()
	fmt.Printf("%q ", modules)
	modules, _ = sess.ListPorts()
	fmt.Printf("%q ", modules)

	fmt.Printf("%q\n", sess)
}
