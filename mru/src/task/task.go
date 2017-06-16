package task

import (
	"condition"
	"log"
	"net/url"
	"routine"
	"rut"
	"taskresult"
)

type Task struct {
	Name          string               `json:"name"`
	PreCondition  *condition.Condition `json:"precondition"`
	Routine       *routine.Routine     `json:"routine"`
	PostCondition *condition.Condition `json:"postcondition"`
	Clear         *routine.Routine     `json:"clear"`
	Description   string               `json:"description"`
}

func (t *Task) Run(db *rut.DB) *taskresult.Result {
	if res := t.CheckPreCondition(db); !res.Success {
		return res
	}

	if res := t.RunMainRoutine(db); !res.Success {
		return res
	}

	if res := t.CheckPostCondition(db); !res.Success {
		return res
	}

	if res := t.RunClearRoutine(db); !res.Success {
		return res
	}

	return &taskresult.Result{Name: t.Name, Description: t.Description, Success: true}
}

func (t *Task) CheckPreCondition(db *rut.DB) *taskresult.Result {
	if err := t.PreCondition.Check(db); err != nil {
		return &taskresult.Result{
			Name:        t.Name,
			Description: t.Description,
			Success:     false,
			Message:     err.Error(),
		}
	}

	return &taskresult.Result{
		Name:        t.Name,
		Description: t.Description,
		Success:     true,
	}
}

func (t *Task) CheckPostCondition(db *rut.DB) *taskresult.Result {
	if err := t.PostCondition.Check(db); err != nil {
		return &taskresult.Result{
			Name:        t.Name,
			Description: t.Description,
			Success:     false,
			Message:     err.Error(),
		}
	}

	return &taskresult.Result{
		Name:        t.Name,
		Description: t.Description,
		Success:     true,
	}

}

func (t *Task) RunMainRoutine(db *rut.DB) *taskresult.Result {
	if err := t.Routine.Run(db); err != nil {
		return &taskresult.Result{
			Name:        t.Name,
			Description: t.Description,
			Success:     false,
			Message:     err.Error(),
		}
	}

	return &taskresult.Result{
		Name:        t.Name,
		Description: t.Description,
		Success:     true,
	}

}

func (t *Task) RunClearRoutine(db *rut.DB) *taskresult.Result {
	if err := t.Clear.Run(db); err != nil {
		return &taskresult.Result{
			Name:        t.Name,
			Description: t.Description,
			Success:     false,
			Message:     err.Error(),
		}
	}

	return &taskresult.Result{
		Name:        t.Name,
		Description: t.Description,
		Success:     true,
	}

}

func IsTaskParamsValid(in url.Values) bool {
	for k, v := range in {
		log.Println(k, "----------->", v, len(v))
		if len(v) == 0 {
			return false
		}
	}
	return true
}
