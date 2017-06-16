package task

import (
	"assertion"
	"condition"
	"errors"
	"fmt"
	"log"
	"net/url"
	"routine"
	"rut"
	"sort"
	"strings"
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
	if result, err := t.CheckPreCondition(db); err != nil {
		return result
	}

	if result, err := t.RunMainRoutine(db); err != nil {
		return result
	}

	if result, err := t.CheckPostCondition(db); err != nil {
		return result
	}

	if result, err := t.RunClearRoutine(db); err != nil {
		return result
	}

	return &taskresult.Result{Name: t.Name, Description: t.Description, Success: true}
}

func (t *Task) CheckPreCondition(db *rut.DB) (*taskresult.Result, error) {
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

func (t *Task) CheckPostCondition(db *rut.DB) (*taskresult.Result, error) {
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

func (t *Task) RunMainRoutine(db *rut.DB) (*taskresult.Result, error) {
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

func (t *Task) RunClearRoutine(db *rut.DB) (*taskresult.Result, error) {
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

func CreateNewTask(in url.Values) (*Task, error) {
	if !IsTaskParamsValid(in) {
		return nil, errors.New("Cannot create new Task due to Invalid input")
	}

	taskname, _ := in["task_name"]
	taskdesc, _ := in["task_description"]
	prec, err := GetPreCondtion(in)
	if err != nil {
		return nil, errors.New("Cannot create new Task: " + err.Error())
	}

	log.Println(prec)

	postc, err := GetPostCondition(in)
	if err != nil {
		return nil, errors.New("Cannot create new Task: " + err.Error())
	}

	log.Println(postc)
	routine, err := GetRoutines(in)
	if err != nil {
		return nil, errors.New("Cannot create new Task: " + err.Error())
	}

	log.Println(routine)
	clear, err := GetClearRoutines(in)
	if err != nil {
		return nil, errors.New("Cannot create new Task: " + err.Error())
	}
	log.Println(clear)

	return &Task{
		PreCondition:  prec,
		PostCondition: postc,
		Routine:       routine,
		Clear:         clear,
		Name:          taskname[0],
		Description:   taskdesc[0],
	}, nil
}

func GetPreCondtion(in url.Values) (*condition.Condition, error) {
	assertmap := make(map[string]*assertion.Assertion, 1)
	assertions := make([]*assertion.Assertion, 0, 1)
	description, _ := in["precondition_description"]
	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "preconditionassertdut" {
				if _, ok := assertmap[fields[1]]; ok {
					log.Println("Same assertion alread exist: ", k)
					continue
				}
				assertmap[fields[1]] = &assertion.Assertion{DUT: v[0], Seq: fields[1]}
			}
		}
	}

	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "preconditionassertmode" {
				assertmap[fields[1]].Mode = v[0]
			} else if fields[0] == "preconditionassertcli" {
				assertmap[fields[1]].Cli = v[0]
			} else if fields[0] == "preconditionassertexpected" {
				assertmap[fields[1]].Expected = v[0]
			}
		}
	}

	if len(assertmap) == 0 {
		return nil, errors.New("Get no preassertion from input request")
	}

	for _, a := range assertmap {
		log.Printf("Find New Assertion: %q in requst", a)
		assertions = append(assertions, a)
	}

	sort.Stable(AssertionSlice(assertions))

	return &condition.Condition{
		Name:        "precondition",
		Description: description[0],
		Assertions:  assertions,
	}, nil
}

func GetPostCondition(in url.Values) (*condition.Condition, error) {
	assertmap := make(map[string]*assertion.Assertion, 1)
	assertions := make([]*assertion.Assertion, 0, 1)
	description, _ := in["postcondition_description"]
	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "postconditionassertdut" {
				if _, ok := assertmap[fields[1]]; ok {
					log.Println("Same assertion alread exist: ", k)
					continue
				}
				assertmap[fields[1]] = &assertion.Assertion{DUT: v[0], Seq: fields[1]}
			}
		}
	}

	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "postconditionassertmode" {
				assertmap[fields[1]].Mode = v[0]
			} else if fields[0] == "postconditionassertcli" {
				assertmap[fields[1]].Cli = v[0]
			} else if fields[0] == "postconditionassertexpected" {
				assertmap[fields[1]].Expected = v[0]
			}
		}
	}

	if len(assertmap) == 0 {
		return nil, errors.New("Get no postassertion from input request")
	}

	for _, a := range assertmap {
		log.Printf("Find New Assertion: %q in requst", a)
		assertions = append(assertions, a)
	}

	sort.Stable(AssertionSlice(assertions))

	return &condition.Condition{
		Name:        "postcondition",
		Description: description[0],
		Assertions:  assertions,
	}, nil
}

func GetRoutines(in url.Values) (*routine.Routine, error) {
	assertmap := make(map[string]*assertion.Assertion, 1)
	assertions := make([]*assertion.Assertion, 0, 1)
	description, _ := in["routine_description"]
	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "routine_command_dut" {
				if _, ok := assertmap[fields[1]]; ok {
					log.Println("Same assertion alread exist: ", k)
					continue
				}
				assertmap[fields[1]] = &assertion.Assertion{DUT: v[0], Seq: fields[1]}
				log.Println("Find New Routines:")
			}
		}
	}

	for k, v := range assertmap {
		log.Println(k, v)
	}

	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			log.Println(k, "+++", v)
			if fields[0] == "routine_command_mode" {
				assertmap[fields[1]].Mode = v[0]
			} else if fields[0] == "routine_command_cli" {
				assertmap[fields[1]].Cli = v[0]
			} else if fields[0] == "routine_command_expected" {
				assertmap[fields[1]].Expected = v[0]
			}
		}
	}

	log.Println("GetRoutines")
	if len(assertmap) == 0 {
		return nil, errors.New("Get no routine assertion from input request")
	}

	for _, a := range assertmap {
		log.Printf("Find New Assertion: %q in requst", a)
		assertions = append(assertions, a)
	}

	sort.Stable(AssertionSlice(assertions))

	return &routine.Routine{
		Name:        "routine",
		Description: description[0],
		Assertions:  assertions,
	}, nil
}

func GetClearRoutines(in url.Values) (*routine.Routine, error) {
	assertmap := make(map[string]*assertion.Assertion, 1)
	assertions := make([]*assertion.Assertion, 0, 1)
	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "routine_command_dut" {
				if _, ok := assertmap[fields[1]]; ok {
					log.Println("Same assertion alread exist: ", k)
					continue
				}
				assertmap[fields[1]] = &assertion.Assertion{DUT: v[0], Seq: fields[1]}
			}
		}
	}

	for k, v := range in {
		if fields := strings.Split(k, "~"); len(fields) == 2 {
			if fields[0] == "clear_command_mode" {
				assertmap[fields[1]].Mode = v[0]
			} else if fields[0] == "clear_command_cli" {
				assertmap[fields[1]].Cli = v[0]
			} else if fields[0] == "clear_command_expected" {
				assertmap[fields[1]].Expected = v[0]
			}
		}
	}

	if len(assertmap) == 0 {
		return nil, errors.New("Get no clear assertion from input request")
	}

	for _, a := range assertmap {
		log.Printf("Find New Assertion: %q in requst", a)
		assertions = append(assertions, a)
	}

	sort.Stable(AssertionSlice(assertions))

	return &routine.Routine{
		Name:       "clear",
		Assertions: assertions,
	}, nil
}

type AssertionSlice []*assertion.Assertion

func (s AssertionSlice) Len() int           { return len(s) }
func (s AssertionSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s AssertionSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }
