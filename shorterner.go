package main

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Mapper struct {
	fullUrl string
	shorternUrl string
}

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

	// Read file
	yamlFile, _ := ioutil.ReadFile("urls.yaml")
	var data map[string]string
	yaml.Unmarshal([]byte(yamlFile), &data)

	switch os.Args[1] {
	case "configure":
		addCommand.Parse(os.Args[2:])
		data[*shortUrl] = *fullUrl
		text, _ := yaml.Marshal(data)
		ioutil.WriteFile("urls.yaml", []byte(text), 0644)
	case "run":
		runCommand.Parse(os.Args[2:])
		fmt.Println("%v", *port)
	default:
		flag.Parse()
		if *listFlag {
			for key, value := range data {
				fmt.Println("Key: ", key, "value: ", value)
			}
		}
		if *deleteFlag != "" {
			fmt.Println("Delete command")
			delete(data, *deleteFlag)
			text, _ := yaml.Marshal(data)
			ioutil.WriteFile("urls.yaml", []byte(text), 0644)
		}
	}
}
