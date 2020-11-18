package main

import (
	"errors"
	"fmt"
	"github.com/ggdream/compile/conf"
	"github.com/ggdream/compile/lib"
	"os"
	"strings"
)



func app() error {
	res, err := lib.GetPlatform()
	if err != nil {
		return err
	}

	targets, err := lib.SelectTarget(res)
	if err != nil {
		return err
	}



	if err := os.RemoveAll("dist"); err != nil {
		return err
	}
	var counter int8
	for _, v := range targets {
		if err := lib.CmdCompile(v); err != nil {
			fmt.Printf("\033[1;31;40mcfailed\033[0m>> %s\n", v)
			counter++
			continue
			//return err
		}
		fmt.Printf("\033[1;46;30msuccess\033[0m>> %s\n", v)
	}

	fmt.Printf("\nAll binarry packages have been packed. There are %d found errors!\n", counter)
	return nil
}

func runCli() error {
	if len(os.Args) == 2 {
		args := strings.Split(os.Args[1], ":")

		switch len(args) {
		case 1:
			conf.NAME = args[0]
		case 2:
			if args[0] != "" {
				conf.NAME = args[0]
			}
			if args[1] != "" {
				conf.VERSION = args[1]
			}
		default:
			return errors.New("args format error")
		}
	} else if len(os.Args) > 2 {
		return errors.New("the wrong number of arguments")
	}

	return app()
}
