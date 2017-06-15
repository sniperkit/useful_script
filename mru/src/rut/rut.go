package rut

import (
	"assert"
	"cline"
	"command"
	"configuration"
	"errors"
	"log"
	"result"
	"routine"
	"script"
	"task"
)

//RUT should be and interface
type RUT struct {
	name string
	cli  *cline.Cli
	L3   bool //We need to have a feature list. For run each case
	L2   bool
}

func NewRUT(conf *configuration.Configuration) (*RUT, error) {
	if conf == nil {
		return nil, errors.New("Invalid config")
	}

	c, err := cline.NewCli(conf)
	if err != nil {
		return nil, errors.New("Cannot create new RUT with: " + err.Error())
	}

	err = c.Init()
	if err != nil {
		return nil, errors.New("Cannot create new RUT with: " + err.Error())
	}

	return &RUT{
		name: conf.DeviceName,
		cli:  c,
	}, nil
}

func (d *RUT) RunTask(t *task.Task) error {
	if err := t.CheckPreCondition(); err != nil {
		log.Println("PreCondition check failed for task: ", t.Name, " with: ", err.Error())
		return errors.New("PreCondition check failed!: " + err.Error())
	}

	for _, r := range t.Routines {
		err := d.RunRoutine(r)
		if r != nil {
			log.Println(err)
		}
		return errors.New("Cannot run task: " + t.Name + " with: " + err.Error())
	}

	if err := t.CheckPostCondition(); err != nil {
		log.Println("PostCondition check failed for task: ", t.Name, " with: ", err.Error())
		return errors.New("PostCondition check failed!: " + err.Error())
	}

	return nil
}

func (d *RUT) RunRoutine(r *routine.Routine) error {
	log.Println("Running Routine: ", r.Name)
	for _, c := range r.CMD {
		_, err := d.RunCommand(c)
		if err != nil {
			log.Println("Error happend when run routine: ", r.Name, " with: ", err.Error())
			return errors.New("Cannot run routine: " + r.Name + " with: " + err.Error())
		}
	}

	for _, a := range r.Assert {
		success := d.Assert(a)
		if !success {
			return errors.New("Assertion failed for routine: " + r.Name + " Message: " + a.String())
		}
	}

	return nil
}

func (d *RUT) Assert(a *assert.Assert) bool {
	result, err := d.RunCommand(a.CMD)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	a.Raw = result

	return a.Do()
}

func (d *RUT) RunCommand(cmd *command.Command) (string, error) {
	data, err := d.cli.RunCommand(cmd)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (r *RUT) RunScript(sc *script.Script) <-chan result.Result {
	res := make(chan result.Result)
	go func(chan<- result.Result) {
		for _, c := range sc.Commands {
			log.Printf("Run command: %v", c)
			data, err := r.cli.RunCommand(c)
			res <- result.Result{
				Command: c.CMD,
				Result:  string(data),
				Err:     err,
			}
		}
		close(res)
	}(res)

	return res
}

func (d *RUT) CreateVlan(id int) error {
	return nil

}

func (d *RUT) DestroyVlan(id int) error {

	return nil
}

func (d *RUT) DestroyAllVlan() error {

	return nil
}

func (d *RUT) CreateVlanInterface(id int, ip string) error {

	return nil
}

func (d *RUT) DestroyVlanInterface(id int) error {

	return nil
}

func (d *RUT) AddIPAddress(ifname, ip string) error {

	return nil
}

func (d *RUT) DelIPAddress(ifname, ip string) error {

	return nil
}

func (d *RUT) AddSecondaryIPAddress(ifname, ip string) error {

	return nil
}

func (d *RUT) DelSecondaryIPAddress(ifname, ip string) error {

	return nil
}

func (d *RUT) DelAllIPAddress(ifname string) error {

	return nil
}

func (d *RUT) CreateOSPFInstance(id, tag string) error {

	return nil
}

func (d *RUT) DestroyOSPFInstance(id string) error {

	return nil
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
