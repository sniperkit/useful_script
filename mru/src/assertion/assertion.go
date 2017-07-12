package assertion

import (
	"command"
	"dsl"
	"fmt"
	"regexp"
	"rut"
	"strings"
	"sync"
)

type Assertion struct {
	DUT        string          `json:"dut"`
	Command    command.Command `json:"command"`
	Expected   string          `json:"expected"`
	UnExpected string
	Raw        string
}

func (a *Assertion) Do(db *rut.DB) (string, bool) {
	dut, ok := db.DB["DUT"+a.DUT]
	if !ok {
		return fmt.Sprintf("DUT %s is not set!", a.DUT), false
	}

	/*  Need more think
	if !dut.IsAlive() {
		return fmt.Sprintf("DUT %s is not alive!", a.DUT), false
	}
	*/

	cmds, err := dsl.Engine.Parse(dut.Device, &a.Command)
	if err != nil {
		return fmt.Sprintf("Parse Instruction error: %s ", err.Error()), false
	}

	var data string
	wg := sync.WaitGroup{}
	for _, c := range cmds {
		wg.Add(1)
		res, err := dut.RunCommand(c)
		if err != nil {
			data += res
			data += fmt.Sprintf("Run Command: %s failed with: %s", a.Command.CMD, err.Error())
		}
		wg.Done()
		data += res
	}
	wg.Wait()

	if IsErrorHappened(data) {
		return data, false
	}

	a.Raw = string(data)
	defer func() { a.Raw = "" }()
	if a.Expected != "" || a.UnExpected != "" {
		msg, ok := a.Verify()
		if !ok {
			dut.GoInitMode() //When error happed, we should go the init mode.
			return fmt.Sprintf("{{{ Assertion  Faild  }}}: on DUT: %s with command: %s. %s", a.DUT, a.Command.CMD, msg), false
		}
		return msg, true
	}

	return fmt.Sprintf("Invlaid assertion, Both expcted and unexpected are empty!"), false
}

func IsErrorHappened(in string) bool {
	if strings.Contains(in, "Invalid") || strings.Contains(in, "invalid") || strings.Contains(in, "INVALID") ||
		strings.Contains(in, "Error") || strings.Contains(in, "error") || strings.Contains(in, "ERROR") ||
		strings.Contains(in, "received SIGSEGV") || strings.Contains(in, "Call backtrace") {
		return true
	}

	return false
}

func (a *Assertion) Verify() (string, bool) {
	var re *regexp.Regexp
	if strings.HasPrefix(a.Expected, "!!") {
		a.UnExpected = strings.TrimLeft(a.Expected, "!!")
	}

	//Unpected has higher priority
	if a.UnExpected != "" {
		re = regexp.MustCompile(a.UnExpected)
		match := re.FindStringSubmatch(a.Raw)
		if match == nil {
			return fmt.Sprintf("{{{ Assertion Success }}}: on DUT: %s with command: %s. UnExpected: %s", a.DUT, a.Command.CMD, a.UnExpected), true
		} else {
			return fmt.Sprintf("<h4>UnExpected: %s, Get: %s</h4>", a.UnExpected, a.Raw), false
		}
	} else {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		if match == nil {
			return fmt.Sprintf("<h4>Expected: %s, Get: %s</h4>", a.Expected, a.Raw), false
		} else {
			return fmt.Sprintf("{{{ Assertion Success }}}: on DUT: %s with command: %s. Expected: %s", a.DUT, a.Command.CMD, a.Expected), true
		}
	}

	return fmt.Sprintf("Invlaid assertion, Both expcted and unexpected are empty!"), false
}

func (a *Assertion) Assert(db *rut.DB) {
	msg, _ := a.Do(db)
	fmt.Println(msg)
}
