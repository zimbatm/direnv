package shell

import (
	"errors"
	"strings"
)

type vim struct{}

// Vim implements Export for the vim editor
//
// http://www.vim.org/
var Vim vim

func (x vim) Name() string {
	return "vim"
}

func (x vim) Hook() (string, error) {
	return "", errors.New("this feature is not supported. Install the direnv.vim plugin instead")
}

func (x vim) Export(e Export) (out string) {
	for key, value := range e {
		if value == nil {
			out += x.unset(key)
		} else {
			out += x.export(key, *value)
		}
	}
	return out
}

func (x vim) export(key, value string) string {
	return "let $" + x.escapeKey(key) + " = " + x.escapeValue(value) + "\n"
}

func (x vim) unset(key string) string {
	return "let $" + x.escapeKey(key) + " = ''\n"
}

// TODO: support keys with special chars or fail
func (x vim) escapeKey(str string) string {
	return str
}

// TODO: Make sure this escaping is valid
func (x vim) escapeValue(str string) string {
	return "'" + strings.Replace(str, "'", "''", -1) + "'"
}
