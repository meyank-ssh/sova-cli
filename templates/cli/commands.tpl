package commands

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ExecuteCommand executes a command with the given arguments
func ExecuteCommand(name string, args ...string) error {
	fmt.Printf("Executing command: %s %s\n", name, strings.Join(args, " "))
	// In a real application, you would execute the command here
	time.Sleep(500 * time.Millisecond) // Simulate execution
	return nil
}

// PrintCommandError prints an error message for a command
func PrintCommandError(command string, err error) {
	fmt.Fprintf(os.Stderr, "Error executing %s: %v\n", command, err)
}

// PrintCommandOutput prints the output of a command
func PrintCommandOutput(command string, output string) {
	fmt.Printf("Output of %s:\n%s\n", command, output)
}

// ConfirmAction asks the user to confirm an action
func ConfirmAction(action string) bool {
	fmt.Printf("Are you sure you want to %s? [y/N] ", action)
	var response string
	fmt.Scanln(&response)
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
} 