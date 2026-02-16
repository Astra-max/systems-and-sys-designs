package lib

import (
	"runtime"
	"os/exec"
	"os"
)

func ProcessInputs(name string) bool {
	return len(name) == 0 || name == ""
}

func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "Windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}