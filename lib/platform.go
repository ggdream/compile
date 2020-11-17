package lib

import (
	"strings"
)

const (
	windows = "windows"
	darwin  = "darwin"
	linux   = "linux"
)

func GetPlatform() ([]string, error) {
	res, err := Cmd("go", "tool", "dist", "list")
	if err != nil {
		return nil, err
	}

	data := strings.Split(res, "\n")

	return filter(data[:len(data)-1]), nil
}

func filter(data []string) (final []string) {
	for _, v := range data {
		if strings.HasPrefix(v, windows) || strings.HasPrefix(v, linux) || strings.HasPrefix(v, darwin) {
			final = append(final, v)
		}
	}
	return
}
