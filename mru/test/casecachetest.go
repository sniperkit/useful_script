package main

import (
	"cache"
	"mcase"
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
}
