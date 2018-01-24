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

//TODO: Add varadic variables to path.Join
func appDir() string {
	switch runtime.GOOS {
	case "darwin":
		return path.Join(os.Getenv("HOME"), "Library", "Application Support")
	case "windows":
		return os.Getenv("APPDATA")
	default:
		return runtime.GOOS
	}
}
