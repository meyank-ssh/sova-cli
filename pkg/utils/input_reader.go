package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type InputReader struct {
	reader *bufio.Reader
}

func NewInputReader() *InputReader {
	return &InputReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (r *InputReader) ReadInput(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

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

func (r *InputReader) ReadInputWithOptions(prompt string, options []string, defaultOption string) (string, error) {
	fmt.Println(prompt)
	for i, option := range options {
		if option == defaultOption {
			color.New(color.FgGreen).Printf("[%d] %s (default)\n", i+1, option)
		} else {
			fmt.Printf("[%d] %s\n", i+1, option)
		}
	}

	fmt.Print("Enter your choice: ")
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	if input == "" {
		return defaultOption, nil
	}

	if num, err := strconv.Atoi(input); err == nil {
		if num >= 1 && num <= len(options) {
			return options[num-1], nil
		}
		return "", fmt.Errorf("invalid option: %d", num)
	}

	for _, option := range options {
		if strings.EqualFold(input, option) {
			return option, nil
		}
	}

	return "", fmt.Errorf("invalid option: %s", input)
}

func (r *InputReader) ConfirmAction(prompt string) (bool, error) {
	fmt.Printf("%s (y/n): ", prompt)
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	input = strings.ToLower(strings.TrimSpace(input))
	return input == "y" || input == "yes", nil
}

func (r *InputReader) ReadInt(prompt string) (int, error) {
	fmt.Print(prompt)
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

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

var DefaultInputReader = NewInputReader()

func ReadInput(prompt string) (string, error) {
	return DefaultInputReader.ReadInput(prompt)
}

func ReadInputWithDefault(prompt, defaultValue string) (string, error) {
	return DefaultInputReader.ReadInputWithDefault(prompt, defaultValue)
}

func ReadInputWithOptions(prompt string, options []string, defaultOption string) (string, error) {
	return DefaultInputReader.ReadInputWithOptions(prompt, options, defaultOption)
}

func ConfirmAction(prompt string) (bool, error) {
	return DefaultInputReader.ConfirmAction(prompt)
}

func ReadInt(prompt string) (int, error) {
	return DefaultInputReader.ReadInt(prompt)
}

func ReadIntWithDefault(prompt string, defaultValue int) (int, error) {
	return DefaultInputReader.ReadIntWithDefault(prompt, defaultValue)
}
