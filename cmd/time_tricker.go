package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	// TODO: refactor if-else-hell
	if len(args) == 3 {
		// start a timer
		fmt.Println("starting timer " + args[2])
	} else if len(args) == 2 {
		// stop the running timer
		if args[1] == "stop" {
			fmt.Println("stopping timer")
		} else {
			fmt.Println("Wrong syntax, did you want to do something else?")
			fmt.Println("Use \"time stop\" to stop the currently running timer.")
		}
	} else {
		fmt.Println("Not the right format")
		fmt.Println("time start <name> to start a timer")
		fmt.Println("time stop to stop the running timer")
	}

}
