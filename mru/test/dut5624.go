package main

import (
	"assert"
	"command"
	"configuration"
	"fmt"
	"routine"
	"rut"
)

func main() {
	base := "SWITCH[A]"
	d, err := rut.NewRUT(&configuration.Configuration{
		DeviceName:     "V8500",
		IP:             "10.71.20.198",
		Port:           "23",
		UserName:       "admin",
		Password:       "",
		EnablePrompt:   ">",
		LoginPrompt:    "login",
		PasswordPrompt: "Password",
		BasePrompt:     base,
		Prompt:         "#",
		ModeDB: map[string]string{
			"login":         "login",
			"password":      "Passowrd:",
			"enable":        base + ">",
			"normal":        base + "#",
			"config":        base + "(config)",
			"config-vlan":   base + "(config-vlan)",
			"config-if":     base + "(config-if[",
			"config-dhcp":   base + "(config-dhcp[",
			"config-router": base + "(config-router)",
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	routines := []*routine.Routine{
		{
			Name: "Create VLAN Interface",
			CMD: []*command.Command{
				{
					Mode: "normal",
					CMD:  "configure terminal",
				},
				{
					Mode: "config",
					CMD:  "vlan database",
				},
				{
					Mode: "config-vlan",
					CMD:  "vlan 1234",
				},
				{
					Mode: "config-vlan",
					CMD:  "exit",
				},
				{
					Mode: "config",
					CMD:  "interface vlan 1234",
				},
				{
					Mode: "config-if",
					CMD:  "ip address 123.1.1.1/24",
				},
				{
					Mode: "config-if",
					CMD:  "no shutdown",
				},
				{
					Mode: "config-if",
					CMD:  "exit",
				},
				{
					Mode: "config",
					CMD:  "interface giga 10/6",
				},
				{
					Mode: "config-if",
					CMD:  "switchport access vlan 1234",
				},
				{
					Mode: "config-if",
					CMD:  "exit",
				},
				{ //Go back to normal mode
					Mode: "config",
					CMD:  "exit",
				},
			},
			Assert: []*assert.Assert{
				{
					Expected: "1234",
					CMD:      &command.Command{Mode: "normal", CMD: "show vlan"},
					Done:     false,
				},
				{
					Expected: `br1234[[:space:]]+up[[:space:]]+up[[:space:]]+123\.1\.1\.1`,
					CMD:      &command.Command{Mode: "normal", CMD: "show ip interface brief"},
					Done:     false,
				},
				{
					Expected: `UP,BROADCAST,RUNNING,MULTICAST`,
					CMD:      &command.Command{Mode: "normal", CMD: "show interface vlan 100"},
					Done:     false,
				},
			},
		},
		{
			Name: "Destroy VLAN Interface",
			CMD: []*command.Command{
				{
					Mode: "normal",
					CMD:  "configure terminal",
				},
				{
					Mode: "config",
					CMD:  "interface vlan 1234",
				},
				{
					Mode: "config-if",
					CMD:  "shutdown",
				},
				{
					Mode: "config-if",
					CMD:  "exit",
				},
				{
					Mode: "config",
					CMD:  "no interface vlan 1234",
				},
				{
					Mode: "config",
					CMD:  "vlan database",
				},
				{
					Mode: "config-vlan",
					CMD:  "no vlan 1234",
				},
				{
					Mode: "config-vlan",
					CMD:  "exit",
				},
				{
					Mode: "config",
					CMD:  "exit",
				},
			},
			Assert: []*assert.Assert{
				{
					UnExpected: "1234",
					CMD:        &command.Command{Mode: "normal", CMD: "show vlan"},
					Done:       false,
				},
				{
					UnExpected: `br1234[[:space:]]+up[[:space:]]+up[[:space:]]+123\.1\.1\.1`,
					CMD:        &command.Command{Mode: "normal", CMD: "show ip interface brief"},
					Done:       false,
				},
				{
					UnExpected: `UP,BROADCAST,RUNNING,MULTICAST`,
					CMD:        &command.Command{Mode: "normal", CMD: "show interface vlan 100"},
					Done:       false,
				},
			},
		},
	}

	for _, r := range routines {
		err := d.RunRoutine(r)
		if err != nil {
			fmt.Println(err)
		}
	}
}
