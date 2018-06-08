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

	fmt.Printf("%q\n", sess)
}
