package libs

import (
	"os"
	"fmt"
)

func getHostName() string {
	host := os.Getenv("HostName")

	if host != "" || len(host) == 0 {
		return "[ simple-ls ]$ "
	}
	hostName := fmt.Sprintf("[~%s~]$ ", host)
	return hostName
}

func TerminalInput(hostName string) {
	fmt.Print(getHostName())
}