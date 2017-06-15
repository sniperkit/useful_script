package result

import (
	"fmt"
)

type Result struct {
	Command string
	Result  string
	Err     error
}

func (r Result) String() string {
	return fmt.Sprintf("Command: %s, Result: %s, error: %s", r.Command, r.Result, r.Err.Error())
}
