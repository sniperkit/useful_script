package main

import (
	"command"
	"context"
	"flag"
	"fmt"
	/*
		"github.com/google/gopacket"
		"github.com/google/gopacket/pcap"
	*/
	"io/ioutil"
	"net"
	"path/filepath"
	"rut"
	"strings"
	"telnet"
	"util"
)

var (
/*
	handle *pcap.Handle
	err    error
*/
)

var IP = flag.String("ip", "10.55.192.210", "IP address of the remote device")
var Port = flag.String("port", "23", "Port to connect")
var Host = flag.String("hostname", "M3000_210", "Host name of the remote device")
var User = flag.String("username", "admin", "Username of the remote device")
var Password = flag.String("password", "", "Passwrod of the remote device")
var Table = flag.String("table", "", "Table name to dump")
var Register = flag.String("register", "", "Register name to dump")
var Command = flag.String("command", "", "Bcm shell command to run")
var Shell = flag.String("shell", "", "shell command to run")
var Config = flag.String("config", "", "Configure file name")
var Protocol = flag.String("protocol", "telnet", "Login protocol(ssh/telnet)")
var Upload = flag.String("upload", "", "upload protocol (ftp/scp)")
var Download = flag.String("download", "", "download protocol (ftp/scp)")
var Server = flag.String("server", "", "server ip")
var ServerUser = flag.String("suser", "", "server user name")
var ServerPassword = flag.String("spass", "", "server password")
var Local = flag.String("local", "", "Local file name(Must use abslute path)")
var Remote = flag.String("remote", "", "Remote file name(Must use abslute path)")
var Tcpdump = flag.String("tcpdump", "", "Run tcpdump on remote device interface")
var TcpdumpCount = flag.String("count", "100", "Packet count to dump")
var TcpdumpFile = flag.String("dfile", "any.pcap", "File name of tcpdump")
var TcpdumpFilter = flag.String("filter", "", "tcpdump packet filter")
var PCAPDecode = flag.String("pcapdecode", "", "Decode a pcap file.")

func main() {
	flag.Parse()

	if *Table == "" && *Register == "" && *Command == "" && *Config == "" && *Shell == "" && *Upload == "" && *Download == "" && *Tcpdump == "" && *PCAPDecode == "" {
		fmt.Println("You must give an operation to run(Dump table/Register, bcmshell command, shell command, config, upload/download, tcpdump")
		return
	}

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

	if *Protocol == "ssh" || *Upload == "ssh" || *Download == "ssh" {
		if *ServerUser == "" || *ServerPassword == "" {
			fmt.Println("Server User/Password is necessary for ssh")
			return
		}
	}

	if *Upload != "" || *Download != "" {
		if *Server == "" {
			fmt.Println("Server IP is necessary for file download/upload!")
			return
		}

		if *ServerUser == "" {
			fmt.Println("Server user name is necessary for file download/upload!")
			return
		}

		if *Local == "" || *Remote == "" {
			fmt.Println("Local/Remote file name is necessary for file Up/Download")
			return
		}

		if !filepath.IsAbs(*Local) || !filepath.IsAbs(*Remote) {
			fmt.Println("Local/Remote file name should use absolute path.")
			return
		}
	}

	dev, err := rut.New(&rut.RUT{
		Name:     "M3000_210",
		Device:   "V5",
		Protocol: *Protocol,
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
		util.SaveToFile(filenameNormalize(*Table)+".txt", []byte(data))
	}

	if *Register != "" {
		data, err = dev.RunCommand(ctx, &command.Command{
			Mode: "shell",
			CMD:  "bcm g " + *Register,
		})
		fmt.Println(data)
		util.SaveToFile(filenameNormalize(*Register)+".txt", []byte(data))
	}

	if *Command != "" {
		data, err = dev.RunCommand(ctx, &command.Command{
			Mode: "shell",
			CMD:  "bcm " + *Command,
		})
		fmt.Println(data)
		util.SaveToFile(filenameNormalize(*Command)+".txt", []byte(data))
	}

	if *Shell != "" {
		data, err = dev.RunCommand(ctx, &command.Command{
			Mode: "shell",
			CMD:  *Shell,
		})
		fmt.Println(data)
		util.SaveToFile(filenameNormalize(*Shell)+".txt", []byte(data))
	}

	if *Config != "" {
		err = Configure(*User, *Password, *IP, *Port, *Config)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	if *Tcpdump != "" {
		err = dev.TCPDUMP(*Tcpdump, *TcpdumpFilter, *TcpdumpFile, *TcpdumpCount)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	if *Upload != "" {
		if *Upload == "ssh" {
			if err = dev.SCPPut(*Local, *Server, *ServerUser, *ServerPassword, filepath.Dir(*Remote)); err != nil {
				fmt.Printf("Cannot upload file: %s to %s with: %s\n", *Local, *IP, err.Error())
				return
			}
		} else {
			if err = dev.FTPPut(*Local, *Server, *ServerUser, *ServerPassword, filepath.Dir(*Remote)); err != nil {
				fmt.Printf("Cannot upload file: %s to %s with: %s\n", *Local, *IP, err.Error())
				return
			}
		}
	}

	if *Download != "" {
		if *Download == "ssh" {
			if err = dev.SCPGet(*Local, *Server, *ServerUser, *ServerPassword, filepath.Dir(*Remote), filepath.Base(*Remote)); err != nil {
				fmt.Printf("Cannot download file: %s from %s with: %s\n", *Remote, *IP, err.Error())
				return
			}
		} else {

			if err = dev.FTPGet(*Local, *Server, *ServerUser, *ServerPassword, filepath.Dir(*Remote), filepath.Base(*Remote)); err != nil {
				fmt.Printf("Cannot download file: %s from %s with: %s\n", *Remote, *IP, err.Error())
				return
			}
		}
	}

	/* Current compile error on 3.14
	if *PCAPDecode != "" {
		if !strings.HasSuffix(*PCAPDecode, ".pcap") {
			fmt.Println("You must give a pcap file for decode")
			return
		}

		handle, err = pcap.OpenOffline(*PCAPDecode)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()

		// Loop through packets in file
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			fmt.Println(packet)
		}
	}
	*/
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

func filenameNormalize(command string) string {
	return strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.TrimSpace(command), "/", "", -1), " ", "", -1), ":", "", -1), ";", "", -1)
}
