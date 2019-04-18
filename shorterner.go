package main

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

type Mapper struct {
	fullUrl string
	shorternUrl string
}

var data map[string]string

func main() {
	
	listFlag := flag.Bool("l", false, "List redirections")
	deleteFlag := flag.String("d", "", "Remove from the list")

	// Add command
	addCommand := flag.NewFlagSet("configure", flag.ExitOnError)
	shortUrl := addCommand.String("a", "", "shortern url")
	fullUrl := addCommand.String("u", "", "full url")

	// Run command
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)
	port := runCommand.Int("p", 8080, "Start a server")

	// Read file
	yamlFile, _ := ioutil.ReadFile("urls.yaml")
	yaml.Unmarshal([]byte(yamlFile), &data)

	switch os.Args[1] {
	case "configure":
		addCommand.Parse(os.Args[2:])
		data[*shortUrl] = *fullUrl
		text, _ := yaml.Marshal(data)
		ioutil.WriteFile("urls.yaml", []byte(text), 0644)
	case "run":
		runCommand.Parse(os.Args[2:])
		http.HandleFunc("/", serverHandler)
		serverPort := fmt.Sprintf(":%d", *port)
		fmt.Println("Server is running at", serverPort)
		if err := http.ListenAndServe(serverPort, nil); err != nil {
			panic(err)
		}

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

func serverHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path
	shortUrl = strings.TrimPrefix(shortUrl, "/")

	if url, ok := data[shortUrl]; ok {
		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		w.Write([]byte("Page not found"))
	}
}
