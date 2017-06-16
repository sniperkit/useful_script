package main

import (
	"rut"
)

func main() {
	rut.New(&rut.RUT{
		Name:     "DUT1",
		Device:   "V8500",
		IP:       "10.71.20.198",
		Port:     "23",
		Username: "admin",
		Password: "",
	})
}
