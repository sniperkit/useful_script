package condition

import (
	"assertion"
)

type Condition struct {
	Name        string
	Assertions  []*assertion.Assertion
	Description string
}
