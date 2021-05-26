package util

import (
	"os"
)

func IsFileExistent(name string) bool {
	if _, err := os.Stat(name); err != nil {
		return !os.IsNotExist(err)
	} else {
		return true
	}
}
