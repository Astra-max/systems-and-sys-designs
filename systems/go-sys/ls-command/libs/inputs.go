package libs

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func TakeCommand() {
	commands := commands()

	if len(commands) >= 1 && commands[0] != "ls" && commands[0] != "clear" {
		fmt.Println("[Error]: Invalid command")
		return
	}

	if commands[0] == "ls" && len(commands) == 1 {
		List(".")
		return
	}

	for _, flag := range commands[1:] {
		switch flag {
		case "-l":
			LongList(".")
		case"-a":
			break
		case "-R":
			break
		case "-t":
			break
		case "-d":
			break
		default:
			fmt.Printf("ls: Invalid option -- %v\n", flag)
			fmt.Println("Try ls man")
			os.Exit(0)
		}
	}
}

func commands() []string {
	TerminalInput()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cmd := scanner.Text()
	command := strings.Split(cmd, " ")
	return command
}

func CommandQue(instructions []string) []string {
	cmdQue := make([]string, 1)

	for _, command := range instructions {
		switch command {
		case "-R":
			getCurr := cmdQue[0]
			cmdQue = append(cmdQue, getCurr)
		}
	}
}