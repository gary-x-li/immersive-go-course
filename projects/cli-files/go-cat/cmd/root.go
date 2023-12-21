package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Execute() {
	helpFlag := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *helpFlag || len(os.Args) != 2 {
		printHelp()
		return
	}

	filepath := os.Args[1]
	content, err := os.ReadFile(filepath)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File or directory does not exist.")
			return
		}
		fmt.Println("Error:", err)
		return
	}

	_, err = os.Stdout.Write(content)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func printHelp() {
	fmt.Println("Usage: go-cat [file-path]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
