package main

import (
	"fmt"
	"github.com/direnv/direnv/shell"
)

var CmdWatch = &Cmd{
	Name:    "watch",
	Desc:    "Adds a path to the list that direnv watches for changes",
	Args:    []string{"[SHELL]", "PATH"},
	Private: false,
	Fn:      watchCommand,
}

func watchCommand(env Env, args []string) (err error) {
	var path, shellName string

	args = args[1:]
	if len(args) < 1 {
		return fmt.Errorf("A path is required to add to the list of watches")
	}
	if len(args) >= 2 {
		shellName = args[0]
		args = args[1:]
	} else {
		shellName = "bash"
	}

	sh := shell.Detect(shellName)

	if sh == nil {
		return fmt.Errorf("Unknown target shell '%s'", shellName)
	}

	path = args[0]

	watches := NewFileTimes()
	watchString, ok := env[DIRENV_WATCHES]
	if ok {
		watches.Unmarshal(watchString)
	}

	watches.Update(path)

	e := make(shell.Export)
	e.Add(DIRENV_WATCHES, watches.Marshal())

	fmt.Printf(sh.Export(e))

	return
}
