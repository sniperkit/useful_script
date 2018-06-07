package main

import (
	//"bufio"
	"fmt"
	//"os"
	"regexp"
	"strings"
	"telnet"
	//"time"
)

var ResultR = regexp.MustCompile("{(?P<result>[[:alnum:][:space:]]*)}")
var BasicResultR = regexp.MustCompile("(?P<status>[[:alnum:]-]+)[[:space:]]+(?P<result>[[:alnum:][:space:]_-{}*]*)")

var Session *telnet.Session
var MGRSession *telnet.Session

var APIs = make(map[string][]string, 10)
var MGRAPIs = make(map[string][]string, 10)

func GetAllObjects() error {
	res, err := Invoke("AgtHelp ListObjects")
	if err != nil {
		return err
	}

	matches := ResultR.FindStringSubmatch(res)
	if len(matches) == 2 {
		res = matches[1]
	}

	objects := strings.Split(res, " ")
	for _, obj := range objects {
		if _, ok := APIs[obj]; !ok {
			APIs[obj] = nil
		} else {
			fmt.Printf("Duplicate object: %s found!\n", obj)
		}
	}

	return nil
}

func GetMethodParameters(object, method string) error {
	cmd := fmt.Sprintf("AgtHelp ListMethodParameters %s %s", object, method)
	res, err := Invoke(cmd)
	if err != nil {
		return err
	}

	fmt.Printf("%20s %30s:   %q\n", object, method, strings.TrimSpace(res))

	return nil
}

func GetAllMethods() error {
	err := GetAllObjects()
	if err != nil {
		return err
	}

	for obj, _ := range APIs {
		cmd := fmt.Sprintf("AgtHelp ListMethods %s", obj)
		res, err := Invoke(cmd)
		if err != nil {
			APIs[obj] = nil
		}

		matches := ResultR.FindStringSubmatch(res)
		if len(matches) == 2 {
			res = matches[1]
		}

		fields := strings.Split(res, " ")
		methods := make([]string, 0, 10)
		for _, field := range fields {
			methods = append(methods, field)
		}

		APIs[obj] = methods
	}

	return nil
}

func MGRGetAllObjects() error {
	res, err := MGRInvoke("AgtHelp ListObjects")
	if err != nil {
		return err
	}

	matches := ResultR.FindStringSubmatch(res)
	if len(matches) == 2 {
		res = matches[1]
	}

	objects := strings.Split(res, " ")
	for _, obj := range objects {
		if _, ok := MGRAPIs[obj]; !ok {
			MGRAPIs[obj] = nil
		} else {
			fmt.Printf("Duplicate object: %s found!\n", obj)
		}
	}

	return nil
}

func MGRGetAllMethods() error {
	err := MGRGetAllObjects()
	if err != nil {
		return err
	}

	for obj, _ := range MGRAPIs {
		cmd := fmt.Sprintf("AgtHelp ListMethods %s", obj)
		res, err := MGRInvoke(cmd)
		if err != nil {
			MGRAPIs[obj] = nil
		}

		matches := ResultR.FindStringSubmatch(res)
		if len(matches) == 2 {
			res = matches[1]
		}

		fields := strings.Split(res, " ")
		methods := make([]string, 0, 10)
		for _, field := range fields {
			methods = append(methods, field)
		}

		MGRAPIs[obj] = methods
	}

	return nil
}

func MGRGetMethodParameters(object, method string) error {
	cmd := fmt.Sprintf("AgtHelp ListMethodParameters %s %s", object, method)
	fmt.Println(cmd)
	res, err := MGRInvoke(cmd)
	if err != nil {
		return err
	}

	fmt.Printf("%20s %30s:   %q\n", object, method, strings.TrimSpace(res))

	return nil
}

func GetAllOpenSessions() ([]string, error) {
	res, err := Invoke("AgtSessionManager ListOpenSessions")
	if err != nil {
		return nil, err
	}

	var sessions = make([]string, 0, 10)
	matches := ResultR.FindStringSubmatch(res)
	if len(matches) == 2 {
		sess := strings.Split(matches[1], " ")
		for _, s := range sess {
			sessions = append(sessions, s)
		}
	}

	return sessions, nil
}

func ConnectToSession(sess string) (*telnet.Session, error) {
	cmd := fmt.Sprintf("AgtSessionManager GetSessionPort %s", sess)
	res, err := Invoke(cmd)
	if err != nil {
		return nil, err
	}

	fmt.Println(res)
	res = strings.TrimSpace(res)

	mgraddr := fmt.Sprintf("%s:%s", "10.71.20.231", res)
	mgrsess, err := telnet.New3(mgraddr)
	if err != nil {
		return nil, err
	}

	return mgrsess, nil
}

func OpenNewSession() (string, error) {
	//AgtOpenSession RouterTester900
	cmd := fmt.Sprintf("AgtSessionManager OpenSession RouterTester900 AGT_SESSION_ONLINE")
	res, err := Invoke(cmd)
	if err != nil {
		return "", err
	}

	return res, nil
}

func main() {
	err := GetAllMethods()
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%q\n", APIs)

	for obj, methods := range APIs {
		for _, method := range methods {
			GetMethodParameters(obj, method)
		}
	}

	id, err := OpenNewSession()
	if err != nil {
		panic(err)
	}

	sess, err := GetAllOpenSessions()
	if err != nil {
		panic(err)
	}

	for _, s := range sess {
		fmt.Println(s)
	}
	MGRSession, err = ConnectToSession(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%q\n", sess)

	err = MGRGetAllMethods()
	if err != nil {
		panic(err)
	}

	for obj, methods := range MGRAPIs {
		for _, method := range methods {
			//	<-time.Tick(time.Duration(time.Millisecond * 1000))
			MGRGetMethodParameters(obj, method)
		}
	}
}

func Run(method ...string) (string, error) {
	if len(method) == 0 {
		return "", fmt.Errorf("You must give the method to run")
	}

	var cmd string
	for _, p := range method {
		cmd += fmt.Sprintf(" %s", p)
	}
	_, err := Session.WriteLine(cmd)
	if err != nil {
		return "", fmt.Errorf("Run %s with error: %s", cmd, err.Error())
	}

	return GetCommandResult()
}

func GetCommandResult() (string, error) {
	var line []byte

	b, err := Session.ReadByte()
	if err != nil {
		return "", fmt.Errorf("Cannot get command result: ", err.Error())
	}
	if b == 'i' {
		line, err = Session.ReadUntilSkip([]string{"\""}, []string{"name"})
		if err != nil {
			return "", fmt.Errorf("Cannot get result with error: %s", err.Error())
		}
		line = []byte(fmt.Sprintf("%c%s", b, line))
	} else {
		sline, err := Session.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("Cannot get result with error: %s", err.Error())
		}
		line = []byte(fmt.Sprintf("%c%s", b, sline))
	}

	return string(line), nil
}

func MGRGetCommandResult() (string, error) {
	var line []byte

	b, err := MGRSession.ReadByte()
	if err != nil {
		return "", fmt.Errorf("Cannot get command result: ", err.Error())
	}
	if b == 'i' {
		line, err = MGRSession.ReadUntilSkip([]string{"\""}, []string{"name"})
		if err != nil {
			return "", fmt.Errorf("Cannot get result with error: %s", err.Error())
		}
		line = []byte(fmt.Sprintf("%c%s", b, line))
	} else {
		sline, err := MGRSession.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("Cannot get result with error: %s", err.Error())
		}
		line = []byte(fmt.Sprintf("%c%s", b, sline))
	}

	return string(line), nil
}

func MGRInvoke(method ...string) (string, error) {
	cmd := fmt.Sprintf("%s ", "invoke")
	for _, p := range method {
		cmd += fmt.Sprintf(" %s", p)
	}
	_, err := MGRSession.WriteLine(cmd)
	if err != nil {
		return "", fmt.Errorf("Run %s with error: %s", cmd, err.Error())
	}

	line, err := MGRGetCommandResult()
	if err != nil {
		return "", fmt.Errorf("Cannot get result of: %s with error: %s", cmd, err.Error())
	}

	res := BasicResultR.FindStringSubmatch(line)
	if len(res) != 3 {
		return "", fmt.Errorf("Run %s with invalid result: %s", cmd, line)
	}

	if res[1] == "0" {
		return res[2], nil
	}

	return "", fmt.Errorf("Run %s with result: (%s: %s)", cmd, res[1], res[2])
}

func Invoke(method ...string) (string, error) {
	cmd := fmt.Sprintf("%s ", "invoke")
	for _, p := range method {
		cmd += fmt.Sprintf(" %s", p)
	}
	_, err := Session.WriteLine(cmd)
	if err != nil {
		return "", fmt.Errorf("Run %s with error: %s", cmd, err.Error())
	}

	line, err := GetCommandResult()
	if err != nil {
		return "", fmt.Errorf("Cannot get result of: %s with error: %s", cmd, err.Error())
	}

	res := BasicResultR.FindStringSubmatch(line)
	if len(res) != 3 {
		return "", fmt.Errorf("Run %s with invalid result: %s", cmd, line)
	}

	if res[1] == "0" {
		return res[2], nil
	}

	return "", fmt.Errorf("Run %s with result: (%s: %s)", cmd, res[1], res[2])
}

func init() {
	sess, err := telnet.New3("10.71.20.231:9001")
	if err != nil {
		panic(err)
	}

	Session = sess

	/*
		go func(r *telnet.Session) {
			for {
				line, err := r.ReadString('\n')
				if err != nil {
					panic(err)
				}
				fmt.Printf("\r> %s", line)
				fmt.Printf("\r>")
			}
		}(sess)

		go func(r *telnet.Session) {
			fmt.Println(">")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				if scanner.Text() == "quit" {
					os.Exit(0)
				}

				line := scanner.Text()
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "run ") {
					line = strings.TrimPrefix(line, "run ")
					if strings.TrimSpace(line) != "" {
						_, err := r.WriteLine(line)
						if err != nil {
							panic(err)
						}
					}
				}
				fmt.Printf(">")
			}
		}(sess)
	*/
}
