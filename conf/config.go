package conf

import (
	"os"
	"path/filepath"
)

var (
	NAME    = ""
	VERSION = "1.0.0"
)

func init() {
	dir, _ := os.Getwd()
	_, NAME = filepath.Split(dir)
}
