package shell

import (
	"encoding/json"
	"errors"
)

type jsonExport struct{}

// JSON implements Export for any editor or tool that supports loading JSON
// as a format.
//
// http://json.org/
var JSON jsonExport

func (s jsonExport) Name() string {
	return "json"
}

func (s jsonExport) Hook() (string, error) {
	return "", errors.New("this feature is not supported")
}

func (s jsonExport) Export(e Export) string {
	out, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		// Should never happen
		panic(err)
	}
	return string(out)
}
