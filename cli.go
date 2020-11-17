package main

import (
	"github.com/ggdream/compile/lib"
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

	for _, v := range targets {
		if err := lib.CmdCompile(v); err != nil {
			return err
		}
	}

	return nil
}
