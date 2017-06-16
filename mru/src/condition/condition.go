package condition

import (
	"assertion"
	"errors"
	"fmt"
	"rut"
)

type Condition struct {
	Name        string
	Assertions  []*assertion.Assertion
	Description string
}

func (cd *Condition) Check(db *rut.DB) error {
	for _, a := range cd.Assertions {
		err := a.Do(db)
		if err != nil {
			return errors.New(fmt.Sprintf("Condition: %s check(%s) failed with: %s", cd.Name, cd.Description, err.Error()))
		}
	}

	return nil
}
