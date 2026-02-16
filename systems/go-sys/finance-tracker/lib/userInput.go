package lib

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

var Man string = "1. add ----> Adds new expense to expense list\n2. list ----> list all available expenses.\n3. exit ----> Exits the program.\n4. Man ----->  displays programs manual.\n5. clear -----> clears screen.\n"

func Message(str string) {
	fmt.Println(str)
}

func UserInput(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(message)
	scanner.Scan()
	return strings.Trim(scanner.Text(), " ")
}