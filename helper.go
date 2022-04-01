package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput() string {
	inputReader := bufio.NewReader(os.Stdin)
	v, _ := inputReader.ReadString('\n')
	v = strings.TrimSpace(v)
	return v
}

// Read input and parse as int, false if user entered non integer
func readInputAsInt() (int, bool) {
	var v int
	var b bool = false
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSpace(input)
	// Check that user input valid selection
	if value, err := strconv.Atoi(input); err == nil {
		v = value
		b = true
	}
	return v, b
}

// Read input and parse as float64, false if user entered non float
func readInputAsFloat() (float64, bool) {
	var v float64
	var b bool = false
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSpace(input)
	// Check that user input valid selection
	if value, err := strconv.ParseFloat(input, 64); err == nil {
		v = value
		b = true
	}
	return v, b
}
