package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// targetSum is the target sum of two inputted numbers for multiplication.
const targetSum = 2020

func main() {
	if isTerminal(os.Stdin) {
		exitUsageError("usage: cat <filename> | day1")
	}

	numbers, err := readNumbers(os.Stdin)
	if err != nil {
		exitError(1, "read: %w", err)
	}

	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			x, y := numbers[i], numbers[j]
			if (x + y) != targetSum {
				continue
			}

			fmt.Fprintf(os.Stdout, "Solution: %d + %d = %d, %d * %d = %d\n", x, y, targetSum, x, y, (x * y))
		}
	}
}

// readNumbers consumes a stream of newline delimited signed integers from the
// given Reader until an io.EOF error is returned by the reader.
func readNumbers(r io.Reader) ([]int64, error) {
	scanner := bufio.NewScanner(r)

	var lineNumber int
	var result []int64

	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			continue
		}

		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNumber, err)
		}

		lineNumber++
		result = append(result, n)
	}

	return result, nil
}

// isTerminal is a crude check on the device type of a file, typically STDIN,
// to check if we are being piped a file or are attached to a character device
// (i.e. a terminal).
func isTerminal(file *os.File) bool {
	stat, _ := file.Stat()

	return (stat.Mode() & os.ModeCharDevice) != 0
}

// exitError prints a formatted message to STDOUT and terminates with the given
// exit code.
func exitError(code int, message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "error: "+message+"\n", args...)
	os.Exit(code)
}

// exitUsageError exits the programme with an error message and a status of 2.
func exitUsageError(message string, args ...interface{}) {
	exitError(2, message, args...)
}
