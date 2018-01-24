package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

func main() {
	res := appDir()
	fmt.Println(res)
}

func appDir() string {
	switch runtime.GOOS {
	case "darwin":
		return path.Join(os.Getenv("HOME"), "Library", "Application Support")
	default:
		return "Error"
	}
}
