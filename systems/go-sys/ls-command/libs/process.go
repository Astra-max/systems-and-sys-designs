package libs

import (
	"os"
	"fmt"
)

func StartProcess() {
	proc, err := os.StartProcess(
		"/bin/ls",
		[]string{"ls", "-l"},
		&os.ProcAttr{
			Files: []*os.File{
				os.Stdin,
			    os.Stdout,
			    os.Stderr,
			},
		},
	)

	if err != nil {
		fmt.Println("Error: creating ls child process")
	}

	proc.Wait()
}