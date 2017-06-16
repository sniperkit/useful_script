package task

import (
	"condition"
	"log"
	"net/url"
	"routine"
	"rut"
	"taskresult"
)

/*
map[preconditionassertdut~0:[DUT1] clear_command_mode~0:[enable] postconditionassertexpected~0:[this must not be empty] clear_command_expected~0:[this can be empty] task_description:[] routine_command_ut~0:[DUT1] postcondition_description:[] feature:[VLAN] group:[L2] routine_command_expected~0:[this can be empty] precondition_description:[] postconditionassertcli~0:[show running-config] routine_command_cli~0:[show running-config] postconditionassertdut~0:[DUT1] clear_command_ut~0:[DUT1] sgroup:[Bridge] postconditionassertmode~0:[enable] continue:[] preconditionassertmode~0:[enable] routine_description:[] routine_command_mode~0:[enable] name:[vLAN DESTROY] preconditionassertexpected~0:[this must not be empty] clear_command_cli~0:[show running-config] preconditionassertcli~0:[show running-config]]
*/

type Task struct {
	Name          string
	PreCondition  *condition.Condition
	Routine       *routine.Routine
	PostCondition *condition.Condition
	Clear         *routine.Routine
	Description   string
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
