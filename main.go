package main

import (
	"flag"
	"fmt"
)

func main() {
	exampleFlag := flag.String("example", "", "This is an example flag")
	flag.Parse()

	fmt.Println("Hi! To get started writing your CLI utility, edit main.go. Try running this with --help!")

	if *exampleFlag != "" {
		fmt.Printf("Value passed to example flag: %s\n", *exampleFlag)
	}
}
