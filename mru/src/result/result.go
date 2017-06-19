package result

import (
	"fmt"
)

type Result struct {
	Command  string
	Result   string
	Err      error
	Group    string
	SubGroup string
	Feature  string
	Case     string
	Message  string
	Pass     bool
}

func (r Result) String() string {
	return fmt.Sprintf("Command: %s, Result: %s, error: %s", r.Command, r.Result, r.Err.Error())
}
