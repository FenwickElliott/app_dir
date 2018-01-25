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
		return path.Join(os.Getenv("HOME"), ".config")
	case "windows":
		return os.Getenv("APPDATA")
	default:
		return runtime.GOOS
	}
}

//Join returns the application storage directory followed by variadic string arguments
func Join(dirs ...string) string {
	res := Is()
	for _, dir := range dirs {
		res = path.Join(res, dir)
	}
	return res
}
