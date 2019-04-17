package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	
	listFlag := flag.Bool("l", false, "List redirections")
	deleteFlag := flag.String("d", "", "Remove from the list")

	// Add command
	addCommand := flag.NewFlagSet("configure", flag.ExitOnError)
	shortUrl := addCommand.String("a", "", "shortern url")
	fullUrl := addCommand.String("u", "", "full url")

	// Run command
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)
	port := addCommand.String("p", "", "add a new url")

	switch os.Args[1] {
	case "configure":
		addCommand.Parse(os.Args[2:])
		fmt.Println("%v", *shortUrl)
		fmt.Println("%v", *fullUrl)
	case "run":
		runCommand.Parse(os.Args[2:])
		fmt.Println("%v", *port)
	default:
		flag.Parse()
		if *listFlag {
			fmt.Println("List of url")
		}
		if *deleteFlag != "" {
			fmt.Println("Delete command")
		}
	}
}
