package main

import (
	"errors"
	"inPacket/logs"
)

func main() {
	lg := new(logs.WriteLogs)
	lg.WriteIntoLogFile(errors.New("Write"))
}
