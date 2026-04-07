package libs

import (
	"os"
	"syscall"
)

func TakeCommand() {
	killSignal := make(chan os.Signal, 1)
	for {}
}