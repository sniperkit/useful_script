package main

import (
	"command"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"rut"
	"strings"
	"telnet"
	"util"
)

var IP = flag.String("ip", "10.55.192.210", "IP address of the remote device")
var Port = flag.String("port", "23", "Port to connect")
var Host = flag.String("hostname", "M3000_210", "Host name of the remote device")
var User = flag.String("username", "admin", "Username of the remote device")
var Password = flag.String("password", "", "Passwrod of the remote device")
var Table = flag.String("table", "", "Table name to dump")
var Register = flag.String("register", "", "Register name to dump")
var Command = flag.String("command", "", "Bcm shell command to run")
var Config = flag.String("config", "", "Configure file name")

func main() {
	flag.Parse()

	ip := net.ParseIP(*IP)
	if ip == nil {
		fmt.Printf("Invalid IP address: %s\n", *IP)
		return
	}

	if *Port == "" {
		fmt.Printf("Invalid port: %s\n", *IP)
		return
	}

	if *Host == "" {
		fmt.Println("Invalid Host name")
		return
	}

	if *User == "" {
		fmt.Println("Invalid user name")
		return
	}

	if *Table == "" && *Register == "" && *Command == "" && *Config == "" {
		fmt.Println("You must give the Table/Register/Config file name or a command to run")
		return
	}

	dev, err := rut.New(&rut.RUT{
		Name:     "M3000_210",
		Device:   "V5",
		IP:       *IP,
		Port:     *Port,
		Username: *User,
		Hostname: *Host,
		Password: *Password,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = dev.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx := context.Background()

	data, err := dev.RunCommand(ctx, &command.Command{
		Mode: "normal",
		CMD:  "q sh -l",
	})

	if *Table != "" {
		data, err = dev.RunCommand(ctx, &command.Command{
			Mode: "shell",
			CMD:  "bcm d all " + *Table,
		})
		fmt.Println(data)
		util.SaveToFile(*Table+".txt", []byte(data))
	}

	if *Register != "" {
		data, err = dev.RunCommand(ctx, &command.Command{
			Mode: "shell",
			CMD:  "bcm g " + *Register,
		})
		fmt.Println(data)
		util.SaveToFile(*Register+".txt", []byte(data))
	}

	if *Command != "" {
		data, err = dev.RunCommand(ctx, &command.Command{
			Mode: "shell",
			CMD:  "bcm " + *Command,
		})
		fmt.Println(data)
		util.SaveToFile(*Command+".txt", []byte(data))
	}

	if *Config != "" {
		err = Configure(*User, *Password, *IP, *Port, *Config)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func Configure(user, pass, ip, port, config string) error {
	c, err := telnet.New(user, pass, ip, port)
	if err != nil {
		return err
	}

	c.WriteLine("enable")
	_, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	c.WriteLine("terminal length 0")
	_, err = c.ReadUntil("#")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	f, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	lines := strings.Split(string(f), "\n")
	for _, l := range lines {
		if strings.TrimSpace(string(l)) == "" {
			continue
		}

		c.WriteLine(l)
		result, err := c.ReadUntil("#")
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		if strings.Contains(string(result), "Invalid input detected at") ||
			strings.Contains(string(result), "VTY configuration is locked by other VTY") {
			fmt.Printf("Run command \"%s\" with error: \n%s\n", string(l), string(result))
			return fmt.Errorf("Run command \"%s\" with error: \n%s\n", string(l), string(result))
		}
		fmt.Println(string(result))
	}

	return nil
}
