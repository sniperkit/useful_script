package main

import (
	"assertion"
	"cache"
	"command"
	"fmt"
	"mcase"
	"rut"
	//"task"
)

func main() {
	db := cache.New("V8500")
	db.Add(&mcase.Case{
		Group:    "L2",
		SubGroup: "Bridge",
		Feature:  "VLAN",
		Name:     "VLAN create",
	})

	db.Add(&mcase.Case{
		Group:    "L2",
		SubGroup: "Bridge",
		Feature:  "VLAN",
		Name:     "VLAN Delete",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv4",
		Feature:  "Static Route",
		Name:     "Static Route Add",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv4",
		Feature:  "Static Route",
		Name:     "Static Route Delete",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv4",
		Feature:  "Static Route",
		Name:     "Static Route Instance",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv4",
		Feature:  "L3 Interface",
		Name:     "Add Interface Address",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv4",
		Feature:  "L3 Interface",
		Name:     "Delete Interface Address",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv6",
		Feature:  "L3 Interface",
		Name:     "Delete Interface Address",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv6",
		Feature:  "L3 Interface",
		Name:     "Add Interface Address",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv6",
		Feature:  "OSPF",
		Name:     "Enable OSPF",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv6",
		Feature:  "OSPF",
		Name:     "Disable OSPF",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv6",
		Feature:  "OSPF",
		Name:     "OSPF Interface Parameters",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv6",
		Feature:  "OSPF",
		Name:     "OSPF Stub Area",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv6",
		Feature:  "OSPF",
		Name:     "OSPF NSSA Area",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv6",
		Feature:  "OSPF",
		Name:     "OSPF Redistribution",
	})

	db.Add(&mcase.Case{
		Group:    "L3",
		SubGroup: "IPv6",
		Feature:  "OSPF",
		Name:     "OSPF Summary",
	})

	cases := db.Dump()
	for _, c := range cases {
		fmt.Printf("%v\n", c)
	}

	c, err := db.Get(&mcase.Case{
		Group:    "L2",
		SubGroup: "Bridge",
		Feature:  "VLAN",
		Name:     "VLAN create",
	})

	if err != nil {
		panic(err)
	}

	r, _ := rut.New(&rut.RUT{
		Device:   "V8500",
		Name:     "DUT1",
		IP:       "10.71.20.198",
		Port:     "23",
		Username: "admin",
		Password: "",
	})
	c.AddRUT(r)

	r, _ = rut.New(&rut.RUT{
		Device:   "V8500",
		Name:     "DUT2",
		IP:       "10.71.20.115",
		Port:     "23",
		Username: "admin",
		Password: "",
	})
	c.AddRUT(r)

	r, _ = rut.New(&rut.RUT{
		Device:   "V5624G",
		Name:     "DUT3",
		IP:       "10.71.20.167",
		Port:     "23",
		Username: "admin",
		Password: "",
	})
	c.AddRUT(r)

	r, _ = rut.New(&rut.RUT{
		Device:   "V5624G",
		Name:     "DUT3",
		IP:       "10.71.20.121",
		Port:     "23",
		Username: "admin",
		Password: "",
	})
	c.AddRUT(r)

	cases = db.Dump()
	for _, c = range cases {
		fmt.Printf("%v\n", c)
	}
	type Command struct {
		Delay  int
		Mode   string
		CMD    string `json:"Command"`
		End    string
		Result string
	}

	c, err = db.Get(&mcase.Case{
		Group:    "L2",
		SubGroup: "Bridge",
		Feature:  "VLAN",
		Name:     "VLAN create",
	})

	if err != nil {
		panic(err)
	}

	c.Init()
	a := assertion.Assertion{
		DUT: "DUT1",
		Command: command.Command{
			Mode: "normal",
			CMD:  "show running-config",
			End:  "#",
		},
		Expected: "122122",
	}

	a.Assert(&c.RUTs)

	a = assertion.Assertion{
		DUT: "DUT2",
		Command: command.Command{
			Mode: "normal",
			CMD:  "show running-config",
			End:  "#",
		},
		Expected: "100",
	}

	a.Assert(&c.RUTs)

	a = assertion.Assertion{
		DUT: "DUT1",
		Command: command.Command{
			Mode: "normal",
			CMD:  "show system",
			End:  "#",
		},
		Expected: "00:d1:cb:00:69:a2",
	}

	a.Assert(&c.RUTs)

	a = assertion.Assertion{
		DUT: "DUT2",
		Command: command.Command{
			Mode: "normal",
			CMD:  "show system",
			End:  "#",
		},
		Expected: "00:d0:cb:00:69:cc",
	}

	a.Assert(&c.RUTs)
}
