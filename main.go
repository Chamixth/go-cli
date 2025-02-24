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

	// Handle "hello" command
	if command == "hello" {
		fmt.Println("Hello, world!")
	} else if command == "ping" {
		// Handle "ping" command
		out, err := exec.Command("ping", "-c", "4", "google.com").Output()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	} else if command == "certbot" {
		// Ensure there is a domain argument
		if len(os.Args) < 3 {
			fmt.Println("Usage: mycli certbot <domain>")
			os.Exit(1)
		}

		// Get the domain from the arguments
		domain := os.Args[2]

		// Execute the certbot command
		fmt.Println("Running certbot for domain:", domain)
		cmd := exec.Command("certbot", "--nginx", "-d", domain)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the command and check for errors
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running certbot:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Unknown command:", command)
	}
}
