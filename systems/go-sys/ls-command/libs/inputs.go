package libs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TakeCommand() {
	commands := commands()
	commands = TrimEmptySpace(commands)

	if len(commands) >= 1 && commands[0] != "ls" && commands[0] != "clear" {
		fmt.Println("[Error]: Invalid command")
		return
	}

	if commands[0] == "ls" && len(commands) == 1 {
		List(".")
		return
	}

	dirEntries := FilterDirectory(commands)

	if len(dirEntries) == 0 {
		dirEntries = []string{"."}
	}

	for _, dir := range dirEntries {
		for _, flag := range commands[1:] {
			switch flag {
			case "-l":
				LongList(dir)
				break
			case "-a":
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
}

func commands() []string {
	TerminalInput()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cmd := strings.TrimSpace(scanner.Text())
	command := strings.Split(cmd, " ")
	return command
}

func TrimEmptySpace(commands []string) []string {
	results := []string{}

	for _, comm := range commands {
		if comm != "" {
			results = append(results, comm)
		}
	}
	return results
}

func CommandQue(instructions []string) []string {
	cmdQue := make([]string, 1)

	for _, command := range instructions {
		switch command {
		case "-R":
			getCurr := cmdQue[0]
			cmdQue = append(cmdQue, getCurr)
		default:
			cmdQue = append(cmdQue, command)
		}
	}
	return cmdQue
}

func FilterDirectory(instructions []string) []string {
	dirList := []string{}

	for i := len(instructions) - 1; i > 0; i-- {
		if !strings.HasPrefix(instructions[i], "-") {
			fileInfo, err := os.Stat(instructions[i])

			if err != nil {
				fmt.Printf("[Error]: %v\n", err)
				return nil
			}

			if fileInfo.IsDir() {
				dirList = append(dirList, instructions[i])
			} else {
				fmt.Printf("[Error]: %v\n", err)
			}
		}
	}
	return dirList
}