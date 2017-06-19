package condition

import (
	"assertion"
	"errors"
	"fmt"
	"rut"
)

type Condition struct {
	Name        string                 `json:"name"`
	Assertions  []*assertion.Assertion `json:"assertions"`
	Description string                 `json:"description"`
}

func (cd *Condition) Check(db *rut.DB) error {
	for _, a := range cd.Assertions {
		msg, ok := a.Do(db)
		if !ok {
			return errors.New(fmt.Sprintf("Condition: %s check(%s) failed with: %s", cd.Name, cd.Description, msg))
		}
	}

	return nil
}
