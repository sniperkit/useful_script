package rut

import (
	"cline"
	"command"
	"configuration"
	"errors"
	"log"
	"result"
	"script"
)

//RUT should be and interface
type RUT struct {
	name string
	cli  *cline.Cli
	L3   bool //We need to have a feature list. For run each case
	L2   bool
}

type DB struct {
	DB map[string]*RUT
}

func (db *DB) GetRUTByName(name string) *RUT {
	if r, ok := db.DB[name]; ok {
		return r
	}
	return nil
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
