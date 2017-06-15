package routine

import (
	"assert"
	"command"
)

type Routine struct {
	Name   string
	CMD    []*command.Command
	Assert []*assert.Assert
}
