package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type fish struct{}

// Fish adds support for the fish shell as a host
var Fish Shell = fish{}

const fishHook = `
function __direnv_export_eval --on-event fish_postexec;
	"{{.SelfPath}}" export fish | source;
end
`

func (sh fish) Hook() (string, error) {
	return fishHook, nil
}

func (sh fish) Export(e ShellExport) (out string) {
	for key, value := range e {
		if value == nil {
			out += sh.unset(key)
		} else {
			out += sh.export(key, *value)
		}
	}
	return out
}

func (sh fish) Dump(env Env) (out string) {
	for key, value := range env {
		out += sh.export(key, value)
	}
	return out
}

func (sh fish) export(key, value string) string {
	if key == "PATH" {
		command := "set -x -g PATH"
		for _, path := range strings.Split(value, ":") {
			command += " " + sh.escape(path)
		}
		return command + ";"
	}
	return "set -x -g " + sh.escape(key) + " " + sh.escape(value) + ";"
}

func (sh fish) unset(key string) string {
	return "set -e -g " + sh.escape(key) + ";"
}

func (sh fish) escape(str string) string {
	in := []byte(str)
	out := "'"

	hex := func(char byte) {
		out += fmt.Sprintf("'\\x%02x'", char)
	}

	backslash := func(char byte) {
		out += string([]byte{BACKSLASH, char})
	}

	escaped := func(str string) {
		out += "'" + str + "'"
	}

	literal := func(char byte) {
		out += string([]byte{char})
	}

	for _, char := range in {
		DecodeRuneInString ¶
		switch {
		case char == TAB:
			escaped(`\t`)
		case char == LF:
			escaped(`\n`)
		case char == CR:
			escaped(`\r`)
		case char <= US:
			hex(char)
		case char == SINGLE_QUOTE:
			backslash(char)
		case char == BACKSLASH:
			backslash(char)
		case char <= TILDA:
			literal(char)
		case char == DEL:
			hex(char)
		default:
			hex(char)
		}
	}

	out += "'"

	return out
}
