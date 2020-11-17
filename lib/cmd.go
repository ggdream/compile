package lib

import (
	"bytes"
	"fmt"
	"github.com/ggdream/compile/conf"
	"os"
	"os/exec"
	"strings"
)

func Cmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var dst bytes.Buffer
	cmd.Stdout = &dst

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return dst.String(), nil
}

func CmdCompile(target string) error {
	info := strings.Split(target, "/")
	if err := os.Setenv("GOOS", info[0]); err != nil {
		return err
	}
	if err := os.Setenv("GOARCH", info[1]); err != nil {

		return err
	}

	path := fmt.Sprintf("dist/%s-%s-%s-%s/%s%s",
		conf.NAME,
		conf.VERSION,
		info[0],
		info[1],
		conf.NAME,
		func() string {
			if info[0] == windows {
				return ".exe"
			}
			return ""
		}(),
	)
	if msg, err := Cmd("go", "build", "-o", path); err != nil {
		fmt.Println(msg, err)
		return err
	}
	return nil
}
