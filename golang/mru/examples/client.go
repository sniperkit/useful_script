package main

import (
	"command"
	"context"
	"fmt"
	"rut"
)

var CTX = context.Background()

func main() {

	dev, err := rut.New(&rut.RUT{
		Name:     "SWITCH",
		Device:   "V5",
		IP:       "10.71.20.102",
		Port:     "22",
		Username: "liwei",
		Protocol: "ssh",
		Hostname: "SWITCH",
		Password: "Lee123!@#",
	})

	if err != nil {
		panic(err)
	}

	err = dev.Init()
	if err != nil {
		panic(err)
	}

	_, err = dev.RunCommand(CTX, &command.Command{
		Mode: "normal",
		CMD:  " config terminal",
	})
	if err != nil {
		fmt.Println(err)
	}

	data, err := dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " show running-config",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " policy mulitcast_225 create",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "config-policy",
		CMD:  " include flow multicast_225",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
