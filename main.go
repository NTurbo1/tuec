package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	
	if len(args) < 2 {
		panic("Not enough args")
	}

	arguments, err := ParseArgs(args);
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		os.Exit(1)
	}

	fmt.Println(arguments)
}
