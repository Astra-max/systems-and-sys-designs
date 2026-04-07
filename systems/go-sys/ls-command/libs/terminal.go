package libs

import (
	"os/user"
	"ls/util"
	"fmt"
)

func getHostName() string {
	host, err := user.Current()

	if err != nil {
		hostName := util.TEST + "[ my-ls ]$ " + util.RESET
		return hostName
	}
	hostName := fmt.Sprintf("%v[~%s~]$%v ", util.TEST,host.Username, util.RESET)
	return hostName
}

func TerminalInput() {
	fmt.Print(getHostName())
}