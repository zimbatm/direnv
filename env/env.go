package env

import (
	"github.com/direnv/direnv/shell"
	"os"
	"strings"
)

type Env map[string]string

// Get returns an Env object from the os.Environ()
func Get() Env {
	return FromEnviron(os.Environ())
}

// FromEnviron transforms a list of "key=value" as gotten from the
// os.Environ() function to the Env object.
//
// Duplicate keys will be lost in the process (in practice it shouldn't
// happend.)
func FromEnviron(pairs []string) Env {
	env := make(Env, len(pairs))

	for _, kv := range pairs {
		kv2 := strings.SplitN(kv, "=", 2)

		key := kv2[0]
		value := kv2[1]

		env[key] = value
	}

	return env
}

// ToEnviron convers an Env object back into a list of "key=value" as used by
// the `os.ProcAttr{}.Env`.
func ToEnviron(env Env) []string {
	pairs := make([]string, len(env))
	index := 0
	for key, value := range env {
		pairs[index] = strings.Join([]string{key, value}, "=")
		index += 1
	}
	return pairs
}

func (env Env) CleanContext() {
	delete(env, DIRENV_DIR)
	delete(env, DIRENV_WATCHES)
	delete(env, DIRENV_DIFF)
}

func LoadEnv(base64env string) (env Env, err error) {
	env = make(Env)
	err = unmarshal(base64env, &env)
	return
}

func (env Env) Copy() Env {
	newEnv := make(Env)

	for key, value := range env {
		newEnv[key] = value
	}

	return newEnv
}

func (env Env) ToShell(sh shell.Shell) string {
	e := make(shell.Export)

	for key, value := range env {
		e.Add(key, value)
	}

	return sh.Export(e)
}

func (env Env) Serialize() string {
	return marshal(env)
}

func (e1 Env) Diff(e2 Env) *Diff {
	return BuildDiff(e1, e2)
}
