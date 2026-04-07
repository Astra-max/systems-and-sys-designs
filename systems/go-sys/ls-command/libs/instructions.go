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

func LongList(directory string) {
	entries, errRead := os.ReadDir(directory)

	if errRead != nil {
		fmt.Printf("[Error]: %v\n", errRead)
		return
	}

	for _, entry := range entries {
		dir, err := entry.Info()
		if err != nil {
			fmt.Printf("[Error]: %v\n", err)
			return
		}
		if dir.IsDir() {
			dirName := dir.Name()
			dirName = util.BOLDBLUE + dirName + util.RESET
			fmt.Println(dir.Mode(), dirName, dir.Size(), dir.ModTime().Format("Jan 01 15:04"))

		} else {
			fmt.Println(dir.Mode(), dir.Name(), dir.Size(), dir.ModTime().Format("Jan 01 15:04"))
		}
	}
}
