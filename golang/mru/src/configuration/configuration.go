package configuration

import (
//"log"
)

type Configuration struct {
	Device         string
	IP             string
	Port           string
	Username       string
	Password       string
	EnablePrompt   string
	LoginPrompt    string
	PasswordPrompt string
	BasePrompt     string
	Prompt         string
	ModeDB         map[string]string
	Name           string
	Hostname       string
	SessionID      string
}

/*

ModeDB: map[string]string{
		"login":         "login",
		"password":      "Passowrd:",
		"enable":        "SWITCH>",
		"normal":        "SWITCH[A]#",
		"config":        "(config)",
		"config-vlan":   "(config-vlan)",
		"config-if":     "(config-if[",
		"config-dhcp":   "(config-dhcp[",
		"config-router": "(config-router)",
		"bridge":        "(bridge)",
		"shell":         "*SWITCH",
		"bcmshell":      "BCM.0>",
	},
*/

var Modes = []string{
	"login",
	"password",
	"enable",
	"normal",
	"config",
	"config-vlan",
	"config-if",
	"config-dhcp",
	"config-router",
	"config-route-map",
	"bridge",
	"shell",
	"bcmshell",
}

var DefaultHostName = "SWITCH"
var DefaultBasePrompt = "SWITCH"
var DefaultEnablePrompt = ">"
var DefaultLoginPrompt = "login"
var DefaultPasswordPrompt = "Password"
var PromptEnd = "#"

func BuildModeDBFromHostNameAndBasePrompt(host, base string) map[string]string {
	if host == "" {
		host = DefaultHostName
	}
	if base == "" {
		base = DefaultBasePrompt
	}

	db := make(map[string]string, len(Modes))
	for _, m := range Modes {
		if m == "login" {
			db[m] = "login"
		} else if m == "enable" {
			db[m] = base + ">"
		} else if m == "passowrd" {
			db[m] = "Password:"
		} else if m == "shell" {
			db[m] = "*" + host + "#"
		} else if m == "bcmshell" {
			db[m] = "BCM.0>"
		} else if m == "normal" {
			db[m] = base + "#"
		} else if m == "config" {
			db[m] = base + "(config)"
		} else if m == "config-vlan" {
			db[m] = base + "(config-vlan"
		} else if m == "config-if" {
			db[m] = base + "(config-if"
		} else if m == "config-dhcp" {
			db[m] = base + "(config-dhcp"
		} else if m == "config-router" {
			db[m] = base + "(config-router"
		} else if m == "bridge" {
			db[m] = base + "(bridge)"
		} else if m == "config-route-map" {
			db[m] = base + "(config-route-map)"
		}
	}

	//log.Printf("Mode DB: %q", db)
	return db
}
