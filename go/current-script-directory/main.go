package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println(GetCurrentScriptDirectory())
}

// GetCurrentScriptDirectory returns the directory of the currently running go script file.
func GetCurrentScriptDirectory() string {
	// NOTE: Replace the 1 with a 0 if you use this code directly, instead of wrapping it in a function.
	_, scriptPath, _, _ := runtime.Caller(1)
	return filepath.Join(scriptPath, "../")
}
