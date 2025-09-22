package main

import (
	"fmt"
	"os"
)

func main() {
	args := PrepArgs(os.Args)

	isHelp, err := CheckForHelp(args)
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		os.Exit(1)
	}
	if isHelp {
		fmt.Println(HelpText)
		return
	}

	arguments, err := ParseArgs(args);
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		os.Exit(1)
	}

	fmt.Println(arguments)
	err = AddRow(arguments)
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		os.Exit(1)
	}
}
