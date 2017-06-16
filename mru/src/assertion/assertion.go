package assertion

import (
	"command"
	"errors"
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

func (a *Assertion) Do(db *rut.DB) error {
	data, err := db.DB[a.DUT].RunCommand(&a.Command)
	if err != nil {
		return errors.New(fmt.Sprintf("Run Command: %s failed with: %s", a.Command.CMD, err.Error()))
	}

	a.Raw = string(data)
	defer func() { a.Raw = "" }()
	if a.Expected != "" || a.UnExpected != "" {
		msg, ok := a.Verify()
		if !ok {
			return errors.New(fmt.Sprintf("Assertion Faild: with command: %s. %s", a.Command.CMD, msg))
		}
		return nil
	}

	return errors.New(fmt.Sprintf("Invlaid assertion, Both expcted and unexpected are empty!"))
}

func (a *Assertion) Verify() (string, bool) {
	var re *regexp.Regexp
	if a.Expected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		if len(match) == 0 {
			return fmt.Sprintf("Expected: %s, Get: %s", a.Expected, a.Raw), false
		} else {
			return "", true
		}
	} else if a.UnExpected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		if len(match) == 0 {
			return "", true
		} else {
			return fmt.Sprintf("UnExpected: %s, Get: %s", a.UnExpected, a.Raw), false
		}
	}

	return fmt.Sprintf("Invlaid assertion, Both expcted and unexpected are empty!"), false
}
