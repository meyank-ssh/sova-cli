package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// InputReader handles user input
type InputReader struct {
	reader *bufio.Reader
}

// NewInputReader creates a new input reader
func NewInputReader() *InputReader {
	return &InputReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

// ReadInput reads a line of input from the user
func (r *InputReader) ReadInput(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// ReadInputWithDefault reads a line of input from the user with a default value
func (r *InputReader) ReadInputWithDefault(prompt, defaultValue string) (string, error) {
	fmt.Printf("%s [%s]: ", prompt, defaultValue)
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue, nil
	}
	return input, nil
}

// ReadInputWithOptions reads input from the user with a list of options
func (r *InputReader) ReadInputWithOptions(prompt string, options []string, defaultOption string) (string, error) {
	// Display options
	fmt.Println(prompt)
	for i, option := range options {
		if option == defaultOption {
			color.New(color.FgGreen).Printf("[%d] %s (default)\n", i+1, option)
		} else {
			fmt.Printf("[%d] %s\n", i+1, option)
		}
	}

	// Read input
	fmt.Print("Enter your choice: ")
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	// Handle empty input (use default)
	if input == "" {
		return defaultOption, nil
	}

	// Try to parse as number
	if num, err := strconv.Atoi(input); err == nil {
		if num >= 1 && num <= len(options) {
			return options[num-1], nil
		}
		return "", fmt.Errorf("invalid option: %d", num)
	}

	// Check if input matches an option
	for _, option := range options {
		if strings.EqualFold(input, option) {
			return option, nil
		}
	}

	return "", fmt.Errorf("invalid option: %s", input)
}

// ConfirmAction asks the user to confirm an action
func (r *InputReader) ConfirmAction(prompt string) (bool, error) {
	fmt.Printf("%s (y/n): ", prompt)
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	input = strings.ToLower(strings.TrimSpace(input))
	return input == "y" || input == "yes", nil
}

// ReadInt reads an integer from the user
func (r *InputReader) ReadInt(prompt string) (int, error) {
	fmt.Print(prompt)
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

// ReadIntWithDefault reads an integer from the user with a default value
func (r *InputReader) ReadIntWithDefault(prompt string, defaultValue int) (int, error) {
	fmt.Printf("%s [%d]: ", prompt, defaultValue)
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue, nil
	}

	return strconv.Atoi(input)
}

// DefaultInputReader is the default input reader instance
var DefaultInputReader = NewInputReader()

// ReadInput is a convenience function that uses the default input reader
func ReadInput(prompt string) (string, error) {
	return DefaultInputReader.ReadInput(prompt)
}

// ReadInputWithDefault is a convenience function that uses the default input reader
func ReadInputWithDefault(prompt, defaultValue string) (string, error) {
	return DefaultInputReader.ReadInputWithDefault(prompt, defaultValue)
}

// ReadInputWithOptions is a convenience function that uses the default input reader
func ReadInputWithOptions(prompt string, options []string, defaultOption string) (string, error) {
	return DefaultInputReader.ReadInputWithOptions(prompt, options, defaultOption)
}

// ConfirmAction is a convenience function that uses the default input reader
func ConfirmAction(prompt string) (bool, error) {
	return DefaultInputReader.ConfirmAction(prompt)
}

// ReadInt is a convenience function that uses the default input reader
func ReadInt(prompt string) (int, error) {
	return DefaultInputReader.ReadInt(prompt)
}

// ReadIntWithDefault is a convenience function that uses the default input reader
func ReadIntWithDefault(prompt string, defaultValue int) (int, error) {
	return DefaultInputReader.ReadIntWithDefault(prompt, defaultValue)
}
