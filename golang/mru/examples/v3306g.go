package main

import (
	"command"
	"context"
	"fmt"
	"rut"
	"util"
)

func main() {
	dev, err := rut.New(&rut.RUT{
		Name:     "SWITCH",
		Device:   "V5",
		IP:       "10.55.192.212",
		Port:     "23",
		Username: "admin",
		Hostname: "SWITCH",
		Password: "Dasan123456",
	})

	if err != nil {
		panic(err)
	}

	dev.Init()

	ctx := context.Background()

	data, err := dev.RunCommand(ctx, &command.Command{
		Mode: "normal",
		CMD:  "configure terminal",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)

	data, err = dev.RunCommand(ctx, &command.Command{
		Mode: "config",
		CMD:  "show running-config",
	})
	fmt.Println(data)

	data, err = dev.RunCommand(ctx, &command.Command{
		Mode: "config",
		CMD:  "do q sh -l",
	})
	fmt.Println(data)

	data, err = dev.RunCommand(ctx, &command.Command{
		Mode: "shell",
		CMD:  " bcm cos bandwidth_show",
	})
	fmt.Println(data)
	util.SaveToFile("cos_bw.txt", []byte(data))

	data, err = dev.RunCommand(ctx, &command.Command{
		Mode: "shell",
		CMD:  " bcm fp show",
	})
	fmt.Println(data)
	util.SaveToFile("fp_show.txt", []byte(data))

	data, err = dev.RunCommand(ctx, &command.Command{
		Mode: "shell",
		CMD:  " bcm d chg fp_tcam",
	})
	fmt.Println(data)
	util.SaveToFile("fp_tcam.txt", []byte(data))

	data, err = dev.RunCommand(ctx, &command.Command{
		Mode: "shell",
		CMD:  " bcm d chg fp_policy_table",
	})
	fmt.Println(data)
	util.SaveToFile("fp_policy_table.txt", []byte(data))
}
