package main

import (
	"cache"
	"fmt"
	"mcase"
	"rut"
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

	c.AddRUT(&rut.RUT{
		Device:   "V8500",
		Name:     "DUT1",
		IP:       "10.71.20.198",
		Port:     "23",
		Username: "admin",
		Password: "",
	})

	c.AddRUT(&rut.RUT{
		Device:   "V8500",
		Name:     "DUT2",
		IP:       "10.71.20.115",
		Port:     "23",
		Username: "admin",
		Password: "",
	})

	c.AddRUT(&rut.RUT{
		Device:   "V5624G",
		Name:     "DUT3",
		IP:       "10.71.20.167",
		Port:     "23",
		Username: "admin",
		Password: "",
	})

	c.AddRUT(&rut.RUT{
		Device:   "V5624G",
		Name:     "DUT3",
		IP:       "10.71.20.121",
		Port:     "23",
		Username: "admin",
		Password: "",
	})

	cases = db.Dump()
	for _, c = range cases {
		fmt.Printf("%v\n", c)
	}
}
