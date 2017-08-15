package main

import (
	"log"
	"os"
	"path/filepath"
	"telnetclient"
	//"util"
	"bufio"
	"bytes"
	"fmt"
	//"unicode"
)

func main() {
	c, err := telnetclient.NewClient("10.71.20.230:10009")
	if err != nil {
		panic(err)
	}

	//buffer := make([]byte, 1024)
	c.SetEcho(false)
	c.SetUnixWriteMode(true)
	c.WriteLine("")
	n, err := c.ReadUntil("login:")
	if err != nil {
		panic(err)
	}

	c.WriteLine("admin")
	n, err = c.ReadUntil("Password:")
	if err != nil {
		panic(err)
	}

	c.WriteLine("")
	n, err = c.ReadUntil(">")
	if err != nil {
		panic(err)
	}

	c.WriteLine("enable")
	n, err = c.ReadUntil("#")
	n, err = c.ReadUntil("#")
	if err != nil {
		panic(err)
	}

	c.WriteLine("terminal length 0")
	n, err = c.ReadUntil("#")
	if err != nil {
		panic(err)
	}

	c.WriteLine("show runnin")
	n, err = c.ReadUntil("#")
	//c.WriteLine("exit")
	//n, err = c.ReadUntil("login")
	fmt.Println(string(n))

	if err != nil {
		panic(err)
	}

	go func() {
		for {
			r, _, e := c.ReadRune()
			if err != nil {
				panic(e)
			}

			fmt.Printf("%c", r)
		}
	}()

	b := bytes.Buffer{}
	for {
		b.WriteRune('a')
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			//fmt.Println(scanner.Text())
			if scanner.Text() == "abcdefg" {
				fmt.Println("We are quiting!")
				os.Exit(0)
			}

			/*
				reader := bytes.NewReader(scanner.Bytes())
				breader := bufio.NewReader(reader)
				_, _, err := breader.ReadRune()
				rn, _, err := breader.ReadRune()
				if err != nil {
					panic(err)
				}
			*/

			//fmt.Println(unicode.IsControl(rn))
			c.WriteLine(scanner.Text())
		}
	}

	/*
		err = util.Wait("fput", "/home/leeweop//workspace/linux/linux-4.13-rc3.tar.gz")
		if err != nil {
			panic(err)
		}
	*/

	abs, err := filepath.Abs("~/workspace/linux/linux-4.13-rc3.tar.gz")
	log.Println(abs, err)

	base := filepath.Base("~/workspace/linux/linux-4.13-rc3.tar.gz")
	log.Println(base)
}

func init() {

}
