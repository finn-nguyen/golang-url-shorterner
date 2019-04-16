package main

import (
	"flag"
	"fmt"
)

func main() {
	list := flag.String("l", "", "List redirections")
	help := flag.String("h", "", "Prints usage info")
	flag.Parse()
	fmt.Println("%v", *list)
	fmt.Println("%v", *help)
}
