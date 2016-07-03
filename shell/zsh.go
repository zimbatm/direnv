package shell

type zsh struct{}

// Zsh implements Export and Hook for the Z shell
//
// http://www.zsh.org/
var Zsh zsh

// ZshHook is the script that will be installed in zsh
const ZshHook = `
_direnv_hook() {
  eval "$(direnv export zsh)";
}
typeset -ag precmd_functions;
if [[ -z ${precmd_functions[(r)_direnv_hook]} ]]; then
  precmd_functions+=_direnv_hook;
fi
`

func (z zsh) Name() string {
	return "zsh"
}

func (z zsh) Hook() (string, error) {
	return ZshHook, nil
}

func (z zsh) Export(e Export) (out string) {
	for key, value := range e {
		if value == nil {
			out += z.unset(key)
		} else {
			out += z.export(key, *value)
		}
	}
	return out
}

func (z zsh) export(key, value string) string {
	return "export " + z.escape(key) + "=" + z.escape(value) + ";"
}

func (z zsh) unset(key string) string {
	return "unset " + z.escape(key) + ";"
}

func (z zsh) escape(str string) string {
	return BashEscape(str)
}
