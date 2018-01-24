package appdir

import (
	"os"
	"path"
	"runtime"
)

//Is returns a string of the appropirate directory on the host machine for application storage
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
