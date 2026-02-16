package lib

import (
	"fmt"
	"os"
	"strconv"
)

type Finance struct {
	name     string
	amount   int
	category string
	list     []Finance
}

func NewFinance() *Finance {
	return new(Finance)
}

func (f *Finance) List() {
	if f.IsEmpty() {
		fmt.Println("Ooops no expenses available")
		return
	}
	for _, val := range f.list {
		fmt.Println(val.name, val.amount, val.category)
	}
}

func (f *Finance) AddExpense(name, category string, amount int) {
	if name == "" {
		fmt.Println("No expense name provided")
	} else if category == "" {
		fmt.Println("expense category not provided")
	} else if amount == 0 {
		fmt.Println("please not amount must be provided")
		return
	}

	f.amount += amount

	f.list = append(f.list, Finance{name: name, amount: amount, category: category})
}

func (f *Finance) IsEmpty() bool {
	return len(f.list) == 0
}

func (f *Finance) Manual() {
	fmt.Println("Finance tracker manual.")
	fmt.Println()
	fmt.Println(Man)
}

func (f *Finance) Controller() {
	f.Manual()
	for {
		command := UserInput("Enter request: ")
		switch command {
		case "":
			Message("Please provide a command")
			break
		case "man":
			f.Manual()
			break
		case "clear":
			Clear()
			f.Manual()
			break
		case "exit":
			os.Exit(0)
		case "total":
			fmt.Printf("Total expenses: %d\n", f.amount)
			break
		case "add":
			name := UserInput("Enter expense name: ")
			errM := ProcessInputs(name)
			if errM {
				Message("please provide exepense name\n")
				break
			}
			category := UserInput("Enter exepense category: ")
			errCategory := ProcessInputs(category)

			if errCategory {
				Message("please provide expense category.\n")
				break
			}
			amount := UserInput("Enter expense amount: ")
			val, err := strconv.Atoi(amount)
			if err != nil {
				fmt.Println("Amount must be a number")
			}
			f.AddExpense(name, category, val)
			break
		case "list":
			f.List()
			break
		default:
			fmt.Println("Unrecognized instruction")
			break
		}
	}
}