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
	"time"
)

//RUT should be and interface
type RUT struct {
	Name     string
	cli      *cline.Cli
	Device   string
	Username string
	Password string
	IP       string
	Port     string
}

type Config struct {
	Index    int    `json:"index"`
	Device   string `json:"device"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"passowrd"`
}

var DefaultConfigurations = configuration.Configuration{
	EnablePrompt:   ">",
	LoginPrompt:    "login",
	PasswordPrompt: "Password",
	Prompt:         "#",
	ModeDB: map[string]string{
		"login":         "login",
		"password":      "Passowrd:",
		"enable":        "SWITCH>",
		"normal":        "SWITCH[A]#",
		"config":        "(config)",
		"config-vlan":   "(config-vlan)",
		"config-if":     "(config-if[",
		"config-dhcp":   "(config-dhcp[",
		"config-router": "(config-router)",
		"shell":         "*SWITCH",
		"bcmshell":      "BCM.0>",
	},
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

func buildDefaultConfiguration(r *RUT) *configuration.Configuration {
	var conf configuration.Configuration
	conf.Name = r.Name
	conf.Username = r.Username
	conf.Password = r.Password
	conf.Device = r.Device
	conf.IP = r.IP
	conf.Port = r.Port

	conf.EnablePrompt = DefaultConfigurations.EnablePrompt
	conf.LoginPrompt = DefaultConfigurations.LoginPrompt
	conf.PasswordPrompt = DefaultConfigurations.PasswordPrompt
	conf.Prompt = DefaultConfigurations.Prompt
	conf.ModeDB = DefaultConfigurations.ModeDB

	return &conf
}

func New(r *RUT) (*RUT, error) {
	return r, nil
}

func (d *RUT) Init() error {
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

func (d *RUT) IsAlive() bool {
	return true
}

func (d *RUT) SetModeDB(db map[string]string) {
	d.cli.SetModeDB(db)
}

func (d *RUT) RunCommand(cmd *command.Command) (string, error) {
	return d.runCommand(cmd)
}

func (d RUT) runCommand(cmd *command.Command) (string, error) {
	<-time.After(time.Second * time.Duration(cmd.Delay))
	data, err := d.cli.RunCommand(cmd)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (r *RUT) RunScript(sc *script.Script) <-chan result.Result {
	res := make(chan result.Result)
	go func(chan<- result.Result) {
		for i := 0; i < sc.Count; i++ {
			for _, c := range sc.Commands {
				<-time.After(time.Second * time.Duration(c.Delay))
				log.Printf("Run command: %v", c)
				data, err := r.cli.RunCommand(c)
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

func GetRUTByConfig(c *Config) (*RUT, error) {
	if !isValidRUTConfig(c) {
		return nil, fmt.Errorf("Invalid config to create RUT: %v", c)
	}
	newrut := &RUT{
		Name:     "DUT" + strconv.Itoa(c.Index),
		Device:   c.Device,
		Username: c.Username,
		Password: c.Password,
		IP:       c.IP,
		Port:     c.Port,
	}

	log.Printf("%q", newrut)

	if err := newrut.Init(); err != nil {
		return nil, fmt.Errorf("Cannot create new DUT with config :%v. Error Message: %s", c, err.Error())
	}

	return newrut, nil
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
