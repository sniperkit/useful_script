package main

import (
	//"cline"
	"command"
	//"configuration"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"rut"
	"strconv"
	"util"
)

var CTX = context.Background()

var IP = flag.String("ip", "10.71.20.10", "IP address of the remote device")
var Host = flag.String("hostname", "SWITCH", "Host name of the remote device")
var User = flag.String("username", "admin", "Username of the remote device")
var Password = flag.String("password", "Dasan123456", "Passwrod of the remote device")
var Start = flag.String("start", "", "start index")
var End = flag.String("end", "", "end index")
var Phase = flag.String("p", "0", "rule stage(0/1)")

func AddRule(dev *rut.RUT, name string) error {
	_, err := dev.RunCommands(CTX, []*command.Command{
		&command.Command{Mode: "config", CMD: " flow " + name + " create"},
		&command.Command{Mode: "config-flow", CMD: " ip any any"},
		&command.Command{Mode: "config-flow", CMD: " apply"},
		&command.Command{Mode: "config-flow", CMD: " exit"},
		&command.Command{Mode: "config", CMD: " policer " + name + " create"},
		&command.Command{Mode: "config-policer", CMD: " counter"},
		&command.Command{Mode: "config-policer", CMD: " apply"},
		&command.Command{Mode: "config-policer", CMD: " exit"},
		&command.Command{Mode: "config", CMD: " policy " + name + " create"},
		&command.Command{Mode: "config-policy", CMD: " include-flow " + name},
		&command.Command{Mode: "config-policy", CMD: " include-policer " + name},
		&command.Command{Mode: "config-policy", CMD: " action match deny"},
		&command.Command{Mode: "config-policy", CMD: " interface-binding port ingress 1-10"},
		&command.Command{Mode: "config-policy", CMD: " apply"},
		&command.Command{Mode: "config-policy", CMD: " exit"},
	})

	return err
}

func DelRule(dev *rut.RUT, name string) error {
	_, err := dev.RunCommands(CTX, []*command.Command{
		&command.Command{Mode: "config", CMD: " no policy " + name},
		&command.Command{Mode: "config", CMD: " no policer " + name},
		&command.Command{Mode: "config", CMD: " no flow " + name},
	})

	return err
}

func DumpTable(dev *rut.RUT, version string) {
	err := os.Remove("FP_TCAM_" + version + ".txt")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	err = os.Remove("FP_POLICY_TABLE_" + version + ".txt")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove("FP_GLOBAL_MASK_TCAM_" + version + ".txt")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove("FP_PORT_FIELD_SEL_" + version + ".txt")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	data, err := dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " do q sh -l",
	})
	if err != nil {
		fmt.Println(err)
	}

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_SLICE_MAP 1 1",
	})
	if err != nil {
		fmt.Println(err)
	}

	util.SaveToFile("FP_SLICE_MAP_"+version+".txt", []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_PORT_FIELD_SEL 0 127",
	})
	if err != nil {
		fmt.Println(err)
	}

	util.SaveToFile("FP_PORT_FIELD_SEL_"+version+".txt", []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_TCAM 0 4095",
	})
	if err != nil {
		fmt.Println(err)
	}

	util.SaveToFile("FP_TCAM_"+version+".txt", []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_GLOBAL_MASK_TCAM 0 4095",
	})
	if err != nil {
		fmt.Println(err)
	}

	util.SaveToFile("FP_GLOBAL_MASK_TCAM_"+version+".txt", []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_POLICY_TABLE 0 4095",
	})

	if err != nil {
		fmt.Println(err)
	}
	util.SaveToFile("FP_POLICY_TABLE_"+version+".txt", []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " exit",
	})

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	flag.Parse()

	ip := net.ParseIP(*IP)
	if ip == nil {
		fmt.Printf("Invalid IP address: %s\n", *IP)
		return
	}

	if *Host == "" {
		fmt.Println("Invalid Host name")
		return
	}

	if *User == "" {
		fmt.Println("Invalid username")
		return
	}

	var si int64
	var ei int64
	if *Start != "" {
		s, err := strconv.ParseInt(*Start, 10, 32)
		if err != nil {
			fmt.Println("Invalid start index to dump")
			return
		}
		if s < 0 {
			fmt.Println("Invalid start index to dump")
			return
		}
		si = s
	} else {
		si = 0
	}

	if *End != "" {
		e, err := strconv.ParseInt(*End, 10, 32)
		if err != nil {
			fmt.Println("Invalid end index to dump")
			return
		}
		if e < 0 {
			fmt.Println("Invalid end index to dump")
			return
		}
		ei = e
	} else {
		ei = 2303
	}

	if si > ei {
		fmt.Println("Start index should be smaller than End index")
		return
	}

	dev, err := rut.New(&rut.RUT{
		Name:     "SWITCH",
		Device:   "V5",
		IP:       *IP,
		Port:     "23",
		Username: *User,
		Hostname: *Host,
		Password: *Password,
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
	fmt.Println(data)

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " show running-config",
	})
	if err != nil {
		fmt.Println(err)
	}

	DelRule(dev, "test33")
	DumpTable(dev, "Before")
	AddRule(dev, "test33")
	DumpTable(dev, "After")

	util.DiffFile("FP_TCAM_Before.txt", "FP_TCAM_After.txt")
	util.DiffFileToText("FP_TCAM_Before.txt", "FP_TCAM_After.txt")
	util.DiffFileToHtml("FP_TCAM_Before.txt", "FP_TCAM_After.txt")
	util.DiffFile("FP_POLICY_TABLE_Before.txt", "FP_POLICY_TABLE_After.txt")
	util.DiffFile("FP_PORT_FIELD_SEL_Before.txt", "FP_PORT_FIELD_SEL_After.txt")
	util.DiffFile("FP_GLOBAL_MASK_TCAM_Before.txt", "FP_GLOBAL_MASK_TCAM_After.txt")

	StartServer()
}

func StartServer() {

	mux := http.NewServeMux()
	//@liwei: This need more analysis.
	mux.HandleFunc("/", Login)
	mux.Handle("/asset/web/", http.FileServer(http.Dir(".")))

	//Add context support
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, _ := ioutil.ReadFile("./" + r.URL.RequestURI())
		w.Write(data)
	}
}
