package main

import (
	//"cline"
	"command"
	//"configuration"
	"context"
	"fmt"
	"rut"
)

var CTX = context.Background()

func main() {
	dev, err := rut.New(&rut.RUT{
		Name:     "V2708M",
		Device:   "V2",
		IP:       "10.71.20.149",
		Port:     "23",
		Username: "admin",
		Hostname: "SWITCH",
		Password: "",
	})

	if err != nil {
		panic(err)
	}

	err = dev.Init()
	if err != nil {
		panic(err)
	}

	data, err := dev.RunCommand(CTX, &command.Command{
		Mode: "normal",
		CMD:  " config terminal",
	})
	if err != nil {
		fmt.Println(err)
	}

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " show running-config",
	})
	if err != nil {
		fmt.Println(err)
	}

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " do q sh -l",
	})
	if err != nil {
		fmt.Println(err)
	}

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " diag",
	})
	if err != nil {
		fmt.Println(err)
	}

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "rtkshell",
		CMD:  " terminal set pager length 0",
	})
	if err != nil {
		fmt.Println(err)
	}

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "rtkshell",
		CMD:  " port dump port all",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	for i := 0; i < 2048; i++ {
		data, err = dev.RunCommand(CTX, &command.Command{
			Mode: "rtkshell",
			CMD:  fmt.Sprintf(" acl get entry phase 0 entry %d", i),
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(data))
	}
}
