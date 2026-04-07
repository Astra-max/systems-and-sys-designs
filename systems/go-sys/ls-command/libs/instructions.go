package libs

import (
	"os"
	"fmt"
	"ls/util"
)


func List(dir string) {
	err := os.Chdir(dir)
	entries, errRead := os.ReadDir(dir)

	if err != nil {
		fmt.Printf("[Error]: %v\n", err)
		return
	}

	if errRead != nil {
		fmt.Printf("[Error]: %v\n", errRead)
		return
	}

	for _, entry := range entries {
		dir,err:= entry.Info()
		if err != nil {
			fmt.Printf("[Error]: %v\n", err)
			return
		}
		if dir.IsDir() {
			dirName := dir.Name()
			dirName = util.BOLDBLUE + dirName + util.RESET
			fmt.Print(dirName + " ")

		} else {
			fmt.Print(dir.Name()+ " ")
		}
	}
	fmt.Println()
}

func LongList(dir string) {
	err := os.Chdir(dir)
	entries, errRead := os.ReadDir(dir)

	if err != nil {
		fmt.Printf("[Error]: %v\n", err)
		return
	}

	if errRead != nil {
		fmt.Printf("[Error]: %v\n", errRead)
		return
	}

	for _, entry := range entries {
		dir,err:= entry.Info()
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
