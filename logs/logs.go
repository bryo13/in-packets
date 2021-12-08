package logs

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

// WriteLogs structures the logs to be written in a file
type WriteLogs struct {
	Lg   *log.Logger
	File *os.File
}

func dir() string {
	home, _ := homedir.Dir()
	location := fmt.Sprintf("%s/%s", home, ".pcktLogs")
	_ = os.Mkdir(location, 0755)
	return location
}

// WriteIntoLogFile takes an error returning a pointer of type Logger. It locates the home dir,opens the log file and appends log data into the log file
func (l *WriteLogs) WriteIntoLogFile(err error) *log.Logger {
	path := dir()
	filePathandName := fmt.Sprintf("%s/%s", path, "pkt-logs.log")

	l.File, _ = os.OpenFile(filePathandName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	l.Lg = log.New(l.File, "pckt ", log.Ldate|log.Ltime)
	log.SetOutput(l.File)
	l.Lg.Println(err)

	defer l.File.Close()

	return l.Lg
}
