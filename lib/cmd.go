package lib

import (
	"bytes"
	"fmt"
	"github.com/ggdream/compile/conf"
	"github.com/ggdream/compile/pack"
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

	path := fmt.Sprintf("dist/_temp/%s_%s/%s%s",
		info[0],
		info[1],
		conf.NAME,
		func() string {
			switch info[0] {
			case windows:
				return ".exe"
			default:
				return ""
			}
		}(),
	)
	fmt.Printf(">>target: %s\n", target)
	if msg, err := Cmd("go", "build", "-o", path); err != nil {
		fmt.Printf("target: %s | err: %s\n", target, msg)
		return err
	}
	fmt.Printf("<<target: %s\n", target)

	var packer pack.Packer
	switch info[0] {
	case windows:
		packer = &pack.ZipPacker{
			DirName: path,
			DstName: fmt.Sprintf("dist/%s-%s-%s.zip", conf.NAME, conf.VERSION, func() string {
				switch info[1] {
				case "386":
					return "win32"
				case "amd64":
					return "win64"
				default:
					return info[1]
				}
			}()),
		}
	default:
		packer = &pack.TgzPacker{
			DirName: path,
			DstName: fmt.Sprintf("dist/%s-%s-%s-%s.tgz", conf.NAME, conf.VERSION, func() string {
				switch info[0] {
				case darwin:
					return "macos"
				default:
					return info[0]
				}
			}(), func() string {
				switch info[1] {
				case "386":
					return "x86"	// i386
				case "amd64":
					return "x86_64"
				default:
					return info[1]
				}
			}()),
		}
	}
	fmt.Printf(">>compress: %s\n", target)
	if err := packer.Pack(); err != nil {
		fmt.Printf("compress: %s | err: %s\n", target, err.Error())
		return err
	}
	fmt.Printf("<<compress: %s\n", target)

	return nil
}
