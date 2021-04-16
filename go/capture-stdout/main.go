package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	result, _ := CaptureStdout(func(w io.Writer) error {
		fmt.Print("Hello, World!")
		return nil
	})

	fmt.Println("Result: " + result)
}

// CaptureStdout captures everything written to stdout from a specific function.
// You can use this method in tests, to validate that your functions writes a string to the terminal.
func CaptureStdout(capture func(w io.Writer) error) (string, error) {
	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return "", fmt.Errorf("could not capture stdout: %w", err)
	}
	os.Stdout = w

	err = capture(w)
	if err != nil {
		return "", fmt.Errorf("error inside capture function while capturing stdout: %w", err)
	}

	err = w.Close()
	if err != nil {
		return "", fmt.Errorf("could not capture stdout: %w", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("could not capture stdout: %w", err)
	}
	os.Stdout = originalStdout
	err = r.Close()
	if err != nil {
		return "", fmt.Errorf("could not capture stdout: %w", err)
	}

	return string(out), nil
}