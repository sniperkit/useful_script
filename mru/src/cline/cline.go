package cline

import (
	"command"
	"configuration"
	"errors"
	"fmt"
	"log"
	"strings"
	"telnetclient"
)

type Cli struct {
	client       *telnetclient.Client
	conf         *configuration.Configuration
	currentMode  string
	modeToPrompt map[string]string
	promptToMode map[string]string
}

func (c *Cli) RunCommand(cmd *command.Command) (result []byte, err error) {
	if cmd.Mode != c.currentMode {
		return nil, errors.New("Error: Command: " + cmd.CMD + " should be run under: " + cmd.Mode + "! But currently we are under: " + c.currentMode + " mode!")
	}

	if strings.HasPrefix(cmd.CMD, "bcm.user.proxy") {
		c.client.WriteLine(cmd.CMD) //For the stupid bcmshell
		cmd.End = "BCM.0>"
	} else if c.currentMode == "bcmshell" && cmd.CMD != "exit" && cmd.CMD != "quit" {
		c.client.WriteLine(cmd.CMD) //For the stupid bcmshell
		cmd.End = "BCM.0>"
	} else {
		c.client.WriteLine(cmd.CMD)
	}
	if cmd.End == "" {
		cmd.End = c.conf.Prompt
	}
	data, err := c.client.ReadUntil(cmd.End)
	if err != nil {
		fmt.Println("Error happend when get login prompt: ", err.Error())
		return nil, err
	}

	if c.IsErrorExist(string(data)) {
		return nil, errors.New("Runn command: " + cmd.CMD + " with error: <<<<<<<<<<<<<<<<" + string(data) + ">>>>>>>>>>>>>>>>")
	}

	old := c.currentMode
	rs := strings.Split(string(data), "\n")
	//log.Println(len(rs))
	//log.Println(c.promptToMode)
	for p, m := range c.promptToMode {
		//log.Println(p, m, rs[len(rs)-1])
		if strings.Contains(rs[len(rs)-1], p) && m != old {
			c.currentMode = m
		}
	}

	//log.Println("Run: ", cmd.CMD, " mode: ", old, " success!")
	if old != c.currentMode {
		log.Println("After run: ", cmd.CMD, " mode switch from: ", old, " to: ", c.currentMode, "!")
	}

	return data, nil
}

func (c *Cli) GoNormalMode() {
	if c.currentMode == "config" ||
		c.currentMode == "config-vlan" ||
		c.currentMode == "config-if" ||
		c.currentMode == "config-dhcp" ||
		c.currentMode == "config-router" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "exit", End: "#"})
	} else if c.currentMode == "shell" ||
		c.currentMode == "bcmshell" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "exit", End: "#"})
		c.GoNormalMode()
	}
}

func (c *Cli) GoShelllMode() {
	if c.currentMode == "config" ||
		c.currentMode == "config-vlan" ||
		c.currentMode == "config-if" ||
		c.currentMode == "config-dhcp" ||
		c.currentMode == "bridge" ||
		c.currentMode == "config-router" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "do q sh -l", End: "#"})
	} else if c.currentMode == "normal" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "q sh -l", End: "#"})
	} else if c.currentMode == "bcmshell" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "exit", End: "#"})
	}
}

func (c *Cli) GoBCMShelllMode() {
	if c.currentMode == "config" ||
		c.currentMode == "config-vlan" ||
		c.currentMode == "config-if" ||
		c.currentMode == "config-dhcp" ||
		c.currentMode == "bridge" ||
		c.currentMode == "config-router" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "do q sh -l", End: "#"})
		c.GoBCMShelllMode()
	} else if c.currentMode == "normal" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "q sh -l", End: "#"})
		c.GoBCMShelllMode()
	} else if c.currentMode == "shell" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "bcm.user.proxy", End: "BCM.0>"})
	}
}

func (c *Cli) GoConfigMode() {
	if c.currentMode == "config-vlan" ||
		c.currentMode == "config-if" ||
		c.currentMode == "config-dhcp" ||
		c.currentMode == "bridge" ||
		c.currentMode == "config-router" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "exit", End: "#"})
	} else if c.currentMode == "shell" ||
		c.currentMode == "bcmshell" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "exit", End: "#"})
		c.GoConfigMode()
	} else if c.currentMode == "normal" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "configure terminal", End: "#"})
	} else if c.currentMode == "enable" {
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "enable", End: "#"})
		c.RunCommand(&command.Command{Mode: c.CurrentMode(), CMD: "configure terminal", End: "#"})
	}
}

func NewCli(conf *configuration.Configuration) (c *Cli, err error) {
	tc, err := telnetclient.NewClient(conf.IP + ":" + conf.Port)
	if err != nil {
		log.Println("error happend when connect to: ", conf.IP, " with: ", err.Error())
		return nil, errors.New("Cannot connect to host")
	}

	return &Cli{
		client:       tc,
		conf:         conf,
		modeToPrompt: make(map[string]string, 1),
		promptToMode: make(map[string]string, 1),
	}, nil
}

func (c *Cli) CurrentMode() string {
	return c.currentMode
}

func (c *Cli) SetModeDB(db map[string]string) {
	c.ClearModeDB()
	for mode, prompt := range db {
		err := c.AddMode(mode, prompt)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (c *Cli) ClearModeDB() {
	c.promptToMode = make(map[string]string, 1)
	c.modeToPrompt = make(map[string]string, 1)
}

func (c *Cli) Init() error {
	for mode, prompt := range c.conf.ModeDB {
		err := c.AddMode(mode, prompt)
		if err != nil {
			fmt.Println(err.Error())
			return errors.New("Register mode db error!")
		}
	}

	err := c.login()
	if err != nil {
		log.Println("Error happened when login: ", err.Error())
		return err
	}

	c.client.WriteLine("enable")
	data, err := c.client.ReadUntil(c.conf.Prompt)
	if err != nil {
		fmt.Println("Error happend when goto enable mode: ", err.Error())
		return err
	}
	fmt.Println(string(data))

	c.client.WriteLine("terminal length 0")
	data, err = c.client.ReadUntil("#")
	if err != nil {
		fmt.Println("Error happend when SetTerminalLength: ", err.Error())
		return err
	}
	c.currentMode = "normal"
	fmt.Println(string(data))

	return nil
}

func (c *Cli) login() error {
	c.client.SetUnixWriteMode(true)
	data, err := c.client.ReadUntil(c.conf.LoginPrompt)
	if err != nil {
		fmt.Println("Error happend when get login: ", err.Error())
		return err
	}
	fmt.Println(string(data))
	c.client.WriteLine(c.conf.Username)
	data, err = c.client.ReadUntil(c.conf.PasswordPrompt)
	if err != nil {
		fmt.Println("Error happend when get login prompt: ", err.Error())
		return err
	}
	fmt.Println(string(data))
	c.client.WriteLine(c.conf.Password)
	data, err = c.client.ReadUntil(c.conf.EnablePrompt)
	if err != nil {
		fmt.Println("Error happend when login: ", err.Error())
		return err
	}
	fmt.Println(string(data))

	return nil
}

func (c *Cli) AddMode(mode, prompt string) error {
	if mode == "" || prompt == "" {
		return errors.New("Invalid input mode: " + mode + " prompt: " + prompt + "!")
	}

	if _, ok := c.modeToPrompt[mode]; ok {
		return errors.New("Same mode: " + mode + " already exist!")
	}

	if _, ok := c.promptToMode[prompt]; ok {
		return errors.New("Same prompt: " + prompt + " already exist!")
	}

	c.modeToPrompt[mode] = prompt
	c.promptToMode[prompt] = mode

	return nil
}

func (c *Cli) GetPromptByMode(mode string) (string, error) {
	if prompt, ok := c.modeToPrompt[mode]; ok {
		return prompt, nil
	}

	return "", errors.New("Mode: " + mode + " does not exist!")
}

func (c *Cli) GetModeByPrompt(prompt string) (string, error) {
	if mode, ok := c.modeToPrompt[prompt]; ok {
		return mode, nil
	}

	return "", errors.New("Prompt: " + prompt + " does not exist!")
}

func (c *Cli) IsErrorExist(result string) bool {
	if strings.Contains(result, "Invalid input detected at") {
		return true
	}
	return false
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
