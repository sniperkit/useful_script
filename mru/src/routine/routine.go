package routine

import (
	"assertion"
	"errors"
	"fmt"
	"rut"
)

type Routine struct {
	Name        string
	Assertions  []*assertion.Assertion
	Description string
}

func (r *Routine) Run(db *rut.DB) error {
	for _, a := range r.Assertions {
		msg, ok := a.Do(db)
		if !ok {
			return errors.New(fmt.Sprintf("Routine: %s failed with: %s", r.Name, msg))
		}
	}

	return nil
}
