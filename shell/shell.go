// Package shell provides hooks and export marshalling facilities for
// various targets (not all of them are shells). The goal is to
// centralize the exports and hooks all in one place, one file per system.
package shell

import (
	"path/filepath"
)

// Shell defined the interface for integration with other tools.
//
// Hook() doesn't always make sense to be implemented, in that case the
// implementation should return an error with an alternative solution.
type Shell interface {
	// Installation script
	Hook() (string, error)
	// Exported diff for the shell to evaluate
	Export(e Export) string
	// Executable name
	Name() string
}

// Export describes what action the target should take with its
// environment.
type Export map[string]*string

// Add tells the target to add the key and value to its environment
func (e Export) Add(key, value string) {
	e[key] = &value
}

// Remove tells the target to remove the key from its environment
func (e Export) Remove(key string) {
	e[key] = nil
}

// Shells contains the list of supported targets
var Shells = []Shell{
	Bash,
	Fish,
	Tcsh,
	Zsh,

	// Only supports exporting
	JSON,
	Vim,
}

// Detect finds a target shell based on the string passed from $0 or other
// means.
//
// The target string should be a relative or absolute path.
func Detect(target string) Shell {
	// $0 starts with "-"
	if target[0] == '-' {
		target = target[1:]
	}
	target = filepath.Base(target)

	for _, shell := range Shells {
		if target == shell.Name() {
			return shell
		}
	}

	return nil
}
