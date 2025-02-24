package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mycli <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	// Example command: if the user passes "hello", print a message
	if command == "hello" {
		fmt.Println("Hello, world!")
	} else if command == "ping" {
		// Example for running a system command
		out, err := exec.Command("ping", "-c", "4", "google.com").Output()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	} else {
		fmt.Println("Unknown command:", command)
	}
}
