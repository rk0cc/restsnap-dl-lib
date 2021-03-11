package misc

import (
	"os"
	"path/filepath"
	"runtime"
)

func getDataDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		if runtime.GOOS == "windows" {
			homeDir = os.Getenv("USERPROFILE")
		} else {
			homeDir = os.Getenv("HOME")
		}
	}
	return filepath.Join(homeDir, ".restsnap")
}

//GetDataPath : Get path inside of the <home>/.restsnap
func GetDataPath(target ...string) string {
	var cloc string = getDataDir()
	for _, subdir := range target {
		cloc = filepath.Join(cloc, subdir)
	}
	return cloc
}
