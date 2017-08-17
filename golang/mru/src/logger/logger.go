package logger

import (
	"context"
	//"fmt"
	"log"
	"os"
	"sync"
)

var loggerLock = sync.Mutex{}
var PrintToTerminal bool = false

func Push(ctx context.Context, message string) {
	sid := ctx.Value("SESSIONID")
	sessionid, _ := sid.(string)

	loggerLock.Lock()
	file, err := os.OpenFile("asset/log/"+sessionid+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Println("cannot Open file: ", sessionid+".log", " ", err.Error())
		return
	}
	file.WriteString(message)
	file.Close()

	//log to detail log file
	full, err := os.OpenFile("asset/log/"+sessionid+"_full.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Println("cannot Open file: ", sessionid+"_full.log", " ", err.Error())
		return
	}
	full.WriteString(message)
	full.Close()

	loggerLock.Unlock()
	if PrintToTerminal {
		fmt.Printf("%s", message)
	}
}
