package main

import (
	"command"
	"errors"
	"fmt"
	"net"
	"regexp"
	"rut"
	"strconv"
	"strings"
)

type RouteEntry struct {
	Index           int64
	Valid           string
	SrcDiscard      string
	RPE             string
	ReservedECMPPTR string
	Reserved0       string
	PRI             string
	NextHopIndex    string
	Length          int64
	Key             net.IP
	Hit             string
	EvenParity      string
	EntryOnly       string
	ECMPPTR         string
	ECMP            string
	DstDiscard      string
	DefaultRoute    string
	Data            string
	Class_ID        string
}

func (re *RouteEntry) String() string {
	if re.Length > 32 {
		return fmt.Sprintf("%s/%d is not a valid IPv4 Address", re.Key, re.Length)
	} else {
		return fmt.Sprintf("[%6d]: %15s/%-2d>>[Nexthop: %s]>>%s>>%s", re.Index, re.Key, re.Length, re.NextHopIndex, re.ECMP, re.ECMPPTR)
	}
}

type FIB struct {
	DB map[string]*RouteEntry
}

func IsValid(es string) bool {
	if es == "" {
		return false
	}

	if !strings.Contains(es, "VALID=1") ||
		!strings.Contains(es, "SRC_DISCARD") ||
		!strings.Contains(es, "NEXT_HOP_INDEX") ||
		!strings.Contains(es, "KEY") ||
		!strings.Contains(es, "LENGTH") ||
		!strings.Contains(es, "ECMP_PTR") {
		return false
	}

	return true
}

func GetIPv4PrefixByHexString(p, plen string) {

}

func FixIPv4Address(s string) net.IP {
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}

	if len(s) == 0 {
		s = "00000000"
	} else if len(s) == 1 {
		s = "0000000" + s
	} else if len(s) == 2 {
		s = "000000" + s
	} else if len(s) == 3 {
		s = "00000" + s
	} else if len(s) == 4 {
		s = "0000" + s
	} else if len(s) == 5 {
		s = "000" + s
	} else if len(s) == 6 {
		s = "00" + s
	} else if len(s) == 7 {
		s = "0" + s
	}

	f1, _ := strconv.ParseInt("0x"+s[:2], 0, 32)
	f2, _ := strconv.ParseInt("0x"+s[2:4], 0, 32)
	f3, _ := strconv.ParseInt("0x"+s[4:6], 0, 32)
	f4, _ := strconv.ParseInt("0x"+s[6:8], 0, 32)

	return net.IPv4(byte(f1), byte(f2), byte(f3), byte(f4))
}

//L3_DEFIP_ALPM_IPV4.*[360456]: <VALID=1,SRC_DISCARD=0,RPE=0,RESERVED_ECMP_PTR=0,RESERVED_0=0,PRI=0,NEXT_HOP_INDEX=1,LENGTH=0x18,KEY=0x46000000,HIT=0,EVEN_PARITY=0,ENTRY_ONLY=0x26118000002,ECMP_PTR=1,ECMP=0,DST_DISCARD=0,DEFAULTROUTE=0,DATA=2,CLASS_ID=0>
func ParseRouteEntryString(es string) (*RouteEntry, error) {
	if !IsValid(es) {
		return nil, errors.New("Invalid input string: " + es)
	}

	var Entry RouteEntry
	if res, err := match(es, getEntryIndex); err != nil {
		panic(err)
	} else {
		index, err := strconv.ParseInt(res, 0, 64)
		if err != nil {
			panic(err)
		}
		Entry.Index = index
	}
	if res, err := match(es, getValidBit); err != nil {
		panic(err)
	} else {
		Entry.Valid = res
	}

	if res, err := match(es, getNextHopIndex); err != nil {
		panic(err)
	} else {
		Entry.NextHopIndex = res
	}

	if res, err := match(es, getLength); err != nil {
		panic(err)
	} else {
		length, err := strconv.ParseInt(res, 0, 64)
		if err != nil {
			panic(err)
		}
		Entry.Length = length
	}

	if res, err := match(es, getKey); err != nil {
		panic(err)
	} else {
		Entry.Key = FixIPv4Address(res)
	}

	if res, err := match(es, getHitBit); err != nil {
		panic(err)
	} else {
		Entry.Hit = res
	}

	if res, err := match(es, getECMPPTR); err != nil {
		panic(err)
	} else {
		Entry.ECMPPTR = res
	}

	if res, err := match(es, getECMP); err != nil {
		panic(err)
	} else {
		Entry.ECMP = res
	}

	if res, err := match(es, getSrcDiscard); err != nil {
		panic(err)
	} else {
		Entry.SrcDiscard = res
	}

	if res, err := match(es, getDstDiscard); err != nil {
		panic(err)
	} else {
		Entry.DstDiscard = res
	}

	if res, err := match(es, getDefaultRoute); err != nil {
		panic(err)
	} else {
		Entry.DefaultRoute = res
	}

	return &Entry, nil
}

func match(s string, r *regexp.Regexp) (string, error) {
	matches := r.FindStringSubmatch(s)
	if len(matches) == 2 {
		return matches[1], nil
	}

	return "", errors.New("Cannot match for string: " + s + " Re: " + r.String())
}

var getEntryIndex = regexp.MustCompile(`[[:word:]_]+\.\*\[(?P<index>[0-9]+)\]`)
var getValidBit = regexp.MustCompile(`VALID=(?P<valid>[0-9]+)`)
var getNextHopIndex = regexp.MustCompile(`NEXT_HOP_INDEX=(?P<nhi>[0x]?[[:alnum:]]+)`)
var getLength = regexp.MustCompile(`LENGTH=(?P<len>[0x]?[[:alnum:]]+)`)
var getKey = regexp.MustCompile(`KEY=(?P<key>[0x]?[[:alnum:]]+)`)
var getHitBit = regexp.MustCompile(`HIT=(?P<hit>[0-9]+)`)
var getECMPPTR = regexp.MustCompile(`ECMP_PTR=(?P<ecmpptr>[0x]?[[:alnum:]]+)`)
var getECMP = regexp.MustCompile(`ECMP=(?P<ecmpptr>[0x]?[[:alnum:]]+)`)
var getSrcDiscard = regexp.MustCompile(`SRC_DISCARD=(?P<srcdis>[0-9]+)`)
var getDstDiscard = regexp.MustCompile(`DST_DISCARD=(?P<dstdis>[0-9]+)`)
var getDefaultRoute = regexp.MustCompile(`DEFAULTROUTE=(?P<default>[0-9]+)`)

func main() {
	dev, err := rut.New(&rut.RUT{
		Name:       "DUT1",
		Device:     "V8500",
		IP:         "10.71.20.115",
		Port:       "23",
		Username:   "admin",
		Hostname:   "V8500",
		Password:   "",
		BasePrompt: "V8500[A]",
	})

	if err != nil {
		panic(err)
	}

	dev.Init()

	res, err := dev.RunCommand(&command.Command{
		Mode: "normal",
		CMD:  "terminal length 0",
	})

	fmt.Println(res, err)

	res, err = dev.RunCommand(&command.Command{
		Mode: "normal",
		CMD:  "configure terminal",
	})

	fmt.Println(res, err)

	res, err = dev.RunCommand(&command.Command{
		Mode: "normal",
		CMD:  "terminal monitor",
	})

	fmt.Println(res, err)

	res, err = dev.RunCommand(&command.Command{
		Mode: "config",
		CMD:  "show ip route",
	})

	fmt.Println(res, err)

	res, err = dev.RunCommand(&command.Command{
		Mode: "config",
		CMD:  "do q sh -l",
	})

	fmt.Println(res, err)

	res, err = dev.RunCommand(&command.Command{
		Mode: "shell",
		CMD:  "ls -al",
	})
	fmt.Println(res, err)

	res, err = dev.RunCommand(&command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 L3_DEFIP_ALPM_IPV4 0 393215  | grep VALID=1",
	})
	fmt.Println(res, err)

	for _, l := range strings.Split(res, "\n") {
		r, _ := ParseRouteEntryString(l)
		fmt.Println(r)
	}
}
