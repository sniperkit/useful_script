package task

import (
	"condition"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
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
	ID            string               `json:"id"`
}

func (t Task) String() string {
	res := fmt.Sprintf("Task: %20s: \n", t.Name)
	res += fmt.Sprintf("         %#v", t.PreCondition)
	res += fmt.Sprintf("         %#v", t.Routine)
	res += fmt.Sprintf("         %#v", t.PreCondition)
	res += fmt.Sprintf("         %#v", t.Clear)
	res += fmt.Sprintf("         %#v", t.Description)
	res += fmt.Sprintf("         %s", t.ID)
	return res
}

func Hash(name []byte) []byte {
	hash := sha1.New()
	return []byte(hex.EncodeToString(hash.Sum([]byte("taskTASK" + string(name)))))
}

func (t *Task) Run(db *rut.DB) *taskresult.Result {
	//This must be at first, I want to do clear work when error happend.
	defer t.RunClearRoutine(db)
	fmt.Printf("[Running Task]: {%s}\n", t.Name)
	if res := t.CheckPreCondition(db); !res.Success {
		return res
	}

	if res := t.RunMainRoutine(db); !res.Success {
		return res
	}

	if res := t.CheckPostCondition(db); !res.Success {
		return res
	}

	return &taskresult.Result{Name: t.Name, Description: t.Description, Success: true}
}

func (t *Task) CheckPreCondition(db *rut.DB) *taskresult.Result {
	for _, r := range db.DB {
		r.GoInitMode()
	}

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
	for _, r := range db.DB {
		r.GoInitMode()
	}

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
	for _, r := range db.DB {
		r.GoInitMode()
	}

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
	for _, r := range db.DB {
		r.GoInitMode()
	}
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

func (t *Task) SetPreCondition(con *condition.Condition) {
	t.PreCondition = con
}

func (t *Task) SetPostCondition(con *condition.Condition) {
	t.PostCondition = con
}

func (t *Task) SetMainRoutine(r *routine.Routine) {
	t.Routine = r
}

func (t *Task) SetClearRoutine(r *routine.Routine) {
	t.Clear = r
}

func (t *Task) SetName(name string) {
	t.Name = name
}

func (t *Task) SetDescription(desc string) {
	t.Description = desc
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
