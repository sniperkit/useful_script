package main

import (
	"configuration"
	"fmt"
	"rut"
)

func main() {
	base := "SWITCH[A]"
	_, err := rut.NewRUT(&configuration.Configuration{
		DeviceName:     "V8500",
		IP:             "10.71.20.198",
		Port:           "23",
		Username:       "admin",
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
}
