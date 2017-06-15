package assertion

import (
	"command"
	"log"
	"regexp"
)

type Assertion struct {
	DUT        string
	Mode       string
	Cli        string
	Command    command.Command
	Seq        string
	Expected   string
	UnExpected string
	Done       bool
	Raw        string
}

func (a *Assertion) Do() bool {
	a.Done = true

	var re *regexp.Regexp
	if a.Expected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		if len(match) == 0 {
			return false
		} else {
			return true
		}
	} else if a.UnExpected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		if len(match) == 0 {
			return true
		} else {
			return false
		}
	}
	log.Println("Invlaid assertion, Both expcted and unexpected are empty!")
	return false
}

func (a *Assertion) String() string {
	var re *regexp.Regexp
	if a.Expected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		log.Println("LEN: +++++++++++++++++++++++: ", len(match))
		if len(match) != 2 {
			return "FAILED: Command: " + a.Command.CMD + " EXPECTED: " + a.Expected + "GET: " + a.Raw
		} else {
			return "SUCCESS: Command: " + a.Command.CMD + " EXPECTED: " + a.Expected
		}
	} else if a.UnExpected != "" {
		re = regexp.MustCompile(a.Expected)
		match := re.FindStringSubmatch(a.Raw)
		log.Println("LEN: +++++++++++++++++++++++: ", len(match))
		if len(match) != 2 {
			return "SUCCESS: Command: " + a.Command.CMD + " UnEXPECTED: " + a.UnExpected
		} else {
			return "FAILED: Command: " + a.Command.CMD + " UnEXPECTED: " + a.UnExpected + "GET: " + a.Raw
		}
	}
	log.Println("Invlaid assertion, Both expcted and unexpected are empty!")
	return "FAILED: Command: " + a.Command.CMD + " EXPECTED: " + a.Expected + "GET: " + a.Raw

}
