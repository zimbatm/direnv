package main

import (
	"fmt"
	"github.com/direnv/direnv/shell"
)

// `direnv hook $0`
var CmdHook = &Cmd{
	Name: "hook",
	Desc: "Used to setup the shell hook",
	Args: []string{"SHELL"},
	Fn: func(env Env, args []string) (err error) {
		var target string

		if len(args) > 1 {
			target = args[1]
		}

		sh := shell.Detect(target)
		if sh == nil {
			return fmt.Errorf("Unknown target shell '%s'", target)
		}

		h, err := sh.Hook()
		if err != nil {
			return err
		}

		fmt.Println(h)

		return
	},
}
