package libs

import (
	"fmt"
	"ls/util"
	"os"
)

func List(directory string) {
	entries, errRead := os.ReadDir(directory)

	if errRead != nil {
		fmt.Printf("[Error]: %v\n", errRead)
		return
	}
	printShort(entries)
}

func LongList(directory string) {
	entries, errRead := os.ReadDir(directory)

	if errRead != nil {
		fmt.Printf("[Error]: %v\n", errRead)
		return
	}
	printLong(entries)
}

func RecursiveList(directory string) {
	fmt.Println(directory + ":")
	entries, err := os.ReadDir(directory)

	if err != nil {
		fmt.Printf("[Error]; %v\n", err)
		return
	}
	printShort(entries)
}

func printLong(entries []os.DirEntry) {
	for _, entry := range entries {
		dir, err := entry.Info()
		if err != nil {
			fmt.Printf("[Error]: %v\n", err)
			return
		}
		if dir.IsDir() {
			dirName := dir.Name()
			dirName = util.BOLDBLUE + dirName + util.RESET
			fmt.Println(dir.Mode(), dirName, dir.Size(), dir.ModTime().Format("Jan 7 15:04"))

		} else {
			fmt.Println(dir.Mode(), dir.Name(), dir.Size(), dir.ModTime().Format("Jan 7 15:04"))
		}
	}
}

func printShort(entries []os.DirEntry) {
	for _, entry := range entries {
		dir, err := entry.Info()
		if err != nil {
			fmt.Printf("[Error]: %v\n", err)
			return
		}
		if dir.IsDir() {
			dirName := dir.Name()
			dirName = util.BOLDBLUE + dirName + util.RESET
			fmt.Print(dirName + " ")

		} else {
			fmt.Print(dir.Name() + " ")
		}
	}
	fmt.Println()
}
