package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogMessage(err error, message ...any) {
	dt := time.Now().UTC().Local()
	msg := fmt.Sprint(dt, "----", message, err)
	fmt.Println(msg)
	log.Println(msg)
}

func PanicIt(err error, message ...any){
	LogMessage(err, message);
	log.Panic(err)
}

func SetLogFile(path string) error {
    logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    log.SetOutput(logFile)
	return nil
}
