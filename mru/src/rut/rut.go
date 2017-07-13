package rut

import (
	"cline"
	"command"
	"configuration"
	"errors"
	"fmt"
	"log"
	"result"
	"script"
	"strconv"
	"strings"
	"time"
)

//RUT should be and interface
type RUT struct {
	Name       string //Name in test case
	cli        *cline.Cli
	Device     string //Device name
	Username   string
	Password   string
	IP         string
	Port       string
	BasePrompt string
	Hostname   string //hostName
}

type Config struct {
	Index      int    `json:"index"`
	Device     string `json:"device"`
	IP         string `json:"ip"`
	Port       string `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"passowrd"`
	BasePrompt string `json:"baseprompt"`
	Hostname   string `json:"hostname"`
}

type DB struct {
	DB map[string]*RUT `json:"-"`
}

func (db *DB) GetRUTByName(name string) *RUT {
	if r, ok := db.DB[name]; ok {
		return r
	}
	return nil
}

func buildDefaultConfiguration(r *RUT) *configuration.Configuration {
	log.Println(r.Hostname, r.Device)
	var conf configuration.Configuration
	conf.Name = r.Name
	conf.Username = r.Username
	conf.Password = r.Password
	conf.Device = r.Device
	conf.Hostname = r.Hostname
	conf.BasePrompt = r.BasePrompt
	conf.IP = r.IP
	conf.Port = r.Port

	conf.EnablePrompt = configuration.DefaultEnablePrompt
	conf.LoginPrompt = configuration.DefaultLoginPrompt
	conf.PasswordPrompt = configuration.DefaultPasswordPrompt
	conf.Prompt = configuration.PromptEnd
	conf.ModeDB = configuration.BuildModeDBFromHostNameAndBasePrompt(r.Hostname, r.BasePrompt)

	log.Printf("%#v", conf)
	return &conf
}

func New(r *RUT) (*RUT, error) {
	return r, nil
}

func (d *RUT) Init() error {
	if d.Device == "V8" {
		d.BasePrompt = d.Hostname + "[A]"
	} else {
		d.BasePrompt = d.Hostname
	}

	conf := buildDefaultConfiguration(d)
	c, err := cline.NewCli(conf)
	if err != nil {
		return errors.New("Cannot create CLI instance: " + err.Error())
	}

	err = c.Init()
	if err != nil {
		return errors.New("Cannot init RUT with: " + err.Error())
	}

	d.cli = c
	return nil
}

func (d *RUT) GoInitMode() {
	d.cli.GoNormalMode()
}

func (d *RUT) SetModeDB(db map[string]string) {
	d.cli.SetModeDB(db)
}

func (d *RUT) RunCommand(cmd *command.Command) (string, error) {
	return d.runCommand(cmd)
}

func (d RUT) runCommand(cmd *command.Command) (string, error) {
	if cmd.Delay != 0 {
		<-time.After(time.Second * time.Duration(cmd.Delay))
	}
	data, err := d.cli.RunCommand(cmd)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (r *RUT) RunScript(sc *script.Script) <-chan result.Result {
	log.Printf("Start Runing Script: %v", sc)
	res := make(chan result.Result)
	go func(chan<- result.Result) {
		for i := 0; i < sc.Count; i++ {
			for _, c := range sc.Commands {
				<-time.After(time.Second * time.Duration(c.Delay))
				log.Printf("Run command: %v", c)
				data, err := r.cli.RunCommand(c)
				log.Println(string(data), err)
				res <- result.Result{
					Command: c.CMD,
					Result:  string(data),
					Err:     err,
				}
			}
			<-time.After(time.Second * time.Duration(sc.Timer))
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

func isValidRUTConfig(c *Config) bool {
	if c.Index < 0 ||
		c.Device == "" ||
		c.IP == "" ||
		c.Port == "" ||
		c.Username == "" {
		return false
	}

	return true
}

func (d *RUT) IsAlive() bool {
	/*
		msg, err := d.cli.GoNormalMode()
		if err != nil {
			log.Println(err, msg)
			return false
		}
	*/

	res, err := d.RunCommand(&command.Command{
		Mode: d.cli.CurrentMode(),
		CMD:  "show running-config",
	})

	if err != nil {
		log.Println(err, res)
		return false
	}

	if strings.Contains(res, d.Hostname) {
		return true
	}
	return false
}

func GetRUTByConfig(c *Config) (*RUT, error) {
	if !isValidRUTConfig(c) {
		return nil, fmt.Errorf("Invalid config to create RUT: %v", c)
	}
	newrut := &RUT{
		Name:       "DUT" + strconv.Itoa(c.Index),
		Device:     c.Device,
		Username:   c.Username,
		Password:   c.Password,
		IP:         c.IP,
		Port:       c.Port,
		Hostname:   c.Hostname,
		BasePrompt: c.BasePrompt,
	}

	log.Printf("%#v", newrut)

	if err := newrut.Init(); err != nil {
		return nil, fmt.Errorf("Cannot create new DUT with config :%v. Error Message: %s", c, err.Error())
	}

	return newrut, nil
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
