package appdir

import (
	"os"
	"path"
	"runtime"
)

func main() {
	// res := appDir()
	// fmt.Println(res)
}

//TODO: Add varadic variables to path.Join
func Is() string {
	switch runtime.GOOS {
	case "darwin":
		return path.Join(os.Getenv("HOME"), "Library", "Application Support")
	case "linux":
		return os.Getenv("XDG_CONFIG_HOME")
	case "windows":
		return os.Getenv("APPDATA")
	default:
		return runtime.GOOS
	}
}
