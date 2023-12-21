package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Execute() {
	helpFlag := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *helpFlag || len(os.Args) > 2 {
		printHelp()
		return
	}

	var path string
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		path = "."
	}

	info, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File or directory does not exist.")
			return
		}
		fmt.Println("Error:", err)
		return
	}

	if info.IsDir() {
		dirEntries, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, dirEntry := range dirEntries {
			fmt.Println(dirEntry.Name())
		}
	} else {
		fmt.Println(info.Name())
	}
}

func printHelp() {
	fmt.Println("Usage: go-ls [path]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
