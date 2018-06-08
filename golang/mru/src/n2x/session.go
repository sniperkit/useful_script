package n2x

import (
	"fmt"
	"strings"
	"telnet"
	"util"
)

type NSession struct {
	ID          string
	Port        string
	Label       string
	Pid         string
	Conn        *telnet.Session
	PortHandler map[string]string
}

func (ns *NSession) ReservePorts(ports ...string) error {
	ns.PortHandler = make(map[string]string, 10)
	if len(ports) == 0 || !ns.isPortsValid(ports...) {
		return fmt.Errorf("You must give the port set to reserver")
	}

	for _, port := range ports {
		cmd := fmt.Sprintf("AgtPortSelector AddPorts")

		cmd += fmt.Sprintf(" %s", port)

		res, err := ns.Invoke(cmd)
		if err != nil {
			return err
		}
		ns.PortHandler[port] = res
	}

	return nil
}

func (ns *NSession) isPortValid(port string) bool {
	port = strings.TrimSpace(port)
	if port != "103/1" && port != "103/2" && port != "103/3" && port != "103/4" &&
		port != "101/1" && port != "101/2" && port != "101/3" && port != "101/4" {
		return false
	}

	return true
}

func (ns *NSession) isPortsValid(ports ...string) bool {
	for _, port := range ports {
		if !ns.isPortValid(port) {
			return false
		}
	}

	return true
}

func (ns *NSession) Invoke(cmds ...string) (string, error) {
	cmd := fmt.Sprintf("%s ", "invoke")
	for _, p := range cmds {
		cmd += fmt.Sprintf(" %s", p)
	}

	util.AppendToFile("n2x_log.txt", []byte(cmd+"\n"))

	_, err := ns.Conn.WriteLine(cmd)
	if err != nil {
		return "", fmt.Errorf("Run %s with error: %s", cmd, err.Error())
	}

	line, err := ns.GetCommandResult()
	if err != nil {
		return "", fmt.Errorf("Cannot get result of: %s with error: %s", cmd, err.Error())
	}

	res := BasicResultR.FindStringSubmatch(line)
	if len(res) != 3 {
		return "", fmt.Errorf("Run %s with invalid result: %s", cmd, line)
	}

	if res[1] == "0" {
		return strings.TrimSpace(res[2]), nil
	}

	return "", fmt.Errorf("Run %s with result: (%s: %s)", cmd, res[1], res[2])

}

func (ns *NSession) GetCommandResult() (string, error) {
	var line []byte

	b, err := ns.Conn.ReadByte()
	if err != nil {
		return "", fmt.Errorf("Cannot get command result: ", err.Error())
	}
	if b == 'i' {
		line, err = ns.Conn.ReadUntilSkip([]string{"\""}, []string{"name"})
		if err != nil {
			return "", fmt.Errorf("Cannot get result with error: %s", err.Error())
		}
		line = []byte(fmt.Sprintf("%c%s", b, line))
	} else {
		sline, err := ns.Conn.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("Cannot get result with error: %s", err.Error())
		}
		line = []byte(fmt.Sprintf("%c%s", b, sline))
	}

	return string(line), nil
}
