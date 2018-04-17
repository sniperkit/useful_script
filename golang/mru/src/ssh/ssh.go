package ssh

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"time"
)

type Session struct {
	session *ssh.Session
	in      chan string
	out     chan string
	w       io.WriteCloser
	r       *bufio.Reader
}

func New(user, password, ip, port string) (*Session, error) {
	sess := new(Session)
	if err := sess.createConnection(user, password, ip+":"+port); err != nil {
		return nil, err
	}
	if err := sess.muxShell(); err != nil {
		return nil, err
	}
	if err := sess.start(); err != nil {
		return nil, err
	}

	res, err := sess.ReadUntil(">")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	return sess, nil
}

func (s *Session) createConnection(user, password, ipPort string) error {
	client, err := ssh.Dial("tcp", ipPort, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 20 * time.Second,
		Config: ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com",
				"arcfour256", "arcfour128", "aes128-cbc", "aes256-cbc", "3des-cbc", "des-cbc",
			},
		},
	})
	if err != nil {
		return err
	}
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	s.session = session
	return nil
}

func (s *Session) muxShell() error {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := s.session.RequestPty("vt100", 80, 40, modes); err != nil {
		return err
	}
	w, err := s.session.StdinPipe()
	if err != nil {
		return err
	}

	s.w = w
	r, err := s.session.StdoutPipe()
	if err != nil {
		return err
	}
	s.r = bufio.NewReader(r)
	return nil
}

func (s *Session) start() error {
	if err := s.session.Shell(); err != nil {
		return err
	}
	return nil
}

func (s *Session) Write(buf []byte) (int, error) {
	return s.w.Write(buf)
}

func (s *Session) WriteLine(data string) (int, error) {
	return s.w.Write([]byte(data + "\n"))
}

func (s *Session) WriteString(data string) (int, error) {
	return s.w.Write([]byte(data))
}

func (s *Session) RunCommand(cmd string) (string, error) {
	s.WriteLine(cmd)
	return "", nil
}

func (s *Session) readUntil(read bool, delims ...string) ([]byte, int, error) {
	if len(delims) == 0 {
		return nil, 0, nil
	}
	p := make([]string, len(delims))
	for i, s := range delims {
		if len(s) == 0 {
			return nil, 0, nil
		}
		p[i] = s
	}
	var line []byte
	for {
		b, err := s.r.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		if read {
			line = append(line, b)
		}
		for i, s := range p {
			if s[0] == b {
				if len(s) == 1 {
					return line, i, nil
				}
				p[i] = s[1:]
			} else {
				p[i] = delims[i]
			}
		}
	}
	panic(nil)
}

func (s *Session) Read(buf []byte) (int, error) {
	return s.r.Read(buf)
}

// SkipBytes works like ReadBytes but skips all read data.
func (s *Session) SkipBytes(delim byte) error {
	for {
		b, err := s.r.ReadByte()
		if err != nil {
			return err
		}
		if b == delim {
			break
		}
	}
	return nil
}

// ReadString works like bufio.ReadString
func (s *Session) ReadString(delim byte) (string, error) {
	bytes, err := s.r.ReadBytes(delim)
	return string(bytes), err
}

// ReadUntilIndex reads from connection until one of delimiters occurs. Returns
// read data and an index of delimiter or error.
func (s *Session) ReadUntilIndex(delims ...string) ([]byte, int, error) {
	line, n, err := s.readUntil(true, delims...)
	//log.Println(string(line))
	return line, n, err
}

// ReadUntil works like ReadUntilIndex but don't return a delimiter index.
func (s *Session) ReadUntil(delims ...string) ([]byte, error) {
	line, _, err := s.readUntil(true, delims...)
	//log.Println(string(line))
	return line, err
}

// SkipUntilIndex works like ReadUntilIndex but skips all read data.
func (s *Session) SkipUntilIndex(delims ...string) (int, error) {
	_, i, err := s.readUntil(false, delims...)
	return i, err
}

// SkipUntil works like ReadUntil but skips all read data.
func (s *Session) SkipUntil(delims ...string) error {
	_, _, err := s.readUntil(false, delims...)
	return err
}

/*
func main() {
    cmds := make([]string, 0)
    cmds = append(cmds, "enable")
    cmds = append(cmds, "show running-config")
    cmds = append(cmds, "show system")
    cmds = append(cmds, "show ip interface brief")

    sess, err := NewSession("admin", "Test123456", "10.71.20.102:22")
    if err != nil {
	panic(err)
    }

    res, err := sess.ReadUntil("SWITCH>")
    if err != nil {
	fmt.Println(err)
    }

    fmt.Println(string(res))

    for _, cmd := range cmds {
	_, err := sess.RunCommand(cmd)
	if err != nil {
	    panic(err)
	}

	res, err := sess.ReadUntil("SWITCH#")
	if err != nil {
	    fmt.Println(err)
	    break
	}

	fmt.Println(string(res))
    }

    time.Sleep(5 * time.Second)
}
*/
