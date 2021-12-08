package privateIP

import (
	"inPacket/logs"
	"net"
	"os"

	"github.com/withmandala/go-log"
)

// checkInterface uses the net package to scan for interfaces found in the node
func checkInterface() []string {
	var interfacesSlice []string
	lg := new(logs.WriteLogs)
	terminalLogger := log.New(os.Stderr)

	interfaces, err := net.Interfaces()
	if err != nil {
		lg.WriteIntoLogFile(err)
		os.Exit(1)
	}

	for _, intaface := range interfaces {
		interfacesSlice = append(interfacesSlice, intaface)
		terminalLogger.Info("Interfaces found: ", intaface)
	}
}
