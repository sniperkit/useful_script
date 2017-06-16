package assertion

import (
	"command"
	"fmt"
	"regexp"
	"rut"
)

type Assertion struct {
	DUT        string
	Command    command.Command
	Expected   string
	UnExpected string
	Raw        string
}

func (a *Assertion) Do(db *rut.DB) (string, bool) {
	data, err := db.DB[a.DUT].RunCommand(&a.Command)
	if err != nil {
		return fmt.Sprintf("Run Command: %s failed with: %s", a.Command.CMD, err.Error()), false
	}

	a.Raw = string(data)
	defer func() { a.Raw = "" }()
	if a.Expected != "" || a.UnExpected != "" {
		msg, ok := a.Verify()
		if !ok {
			return fmt.Sprintf("{{{ Assertion  Faild  }}}: on DUT: %s with command: %s. %s", a.DUT, a.Command.CMD, msg), false
		}
		return msg, true
	}

	return fmt.Sprintf("Invlaid assertion, Both expcted and unexpected are empty!"), false
}

func (a *Assertion) Verify() (string, bool) {
	var re *regexp.Regexp
	if a.Expected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		if match == nil {
			return fmt.Sprintf("Expected: %s, Get: %s", a.Expected, a.Raw), false
		} else {
			return fmt.Sprintf("{{{ Assertion Success }}}: on DUT: %s with command: %s. Expected: %s", a.DUT, a.Command.CMD, a.Expected), true
		}
	} else if a.UnExpected != "" {
		re = regexp.MustCompile(a.UnExpected)
		match := re.FindStringSubmatch(a.Raw)
		if match == nil {
			return fmt.Sprintf("{{{ Assertion Success }}}: on DUT: %s with command: %s. UnExpected: %s", a.DUT, a.Command.CMD, a.UnExpected), true
		} else {
			return fmt.Sprintf("UnExpected: %s, Get: %s", a.UnExpected, a.Raw), false
		}
	}

	return fmt.Sprintf("Invlaid assertion, Both expcted and unexpected are empty!"), false
}

func (a *Assertion) Assert(db *rut.DB) {
	msg, _ := a.Do(db)
	fmt.Println(msg)
}
