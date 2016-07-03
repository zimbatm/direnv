package shell

import (
	"testing"
)

func TestBashEscape(t *testing.T) {
	assertEqual(t, `''`, BashEscape(""))
	assertEqual(t, `$'escape\'quote'`, BashEscape("escape'quote"))
	assertEqual(t, `$'foo\r\n\tbar'`, BashEscape("foo\r\n\tbar"))
	assertEqual(t, `$'foo bar'`, BashEscape("foo bar"))
	assertEqual(t, `$'\xc3\xa9'`, BashEscape("é"))
}

func TestShellDetection(t *testing.T) {
	assertNotNil(t, Detect("-bash"))
	assertNotNil(t, Detect("-/bin/bash"))
	assertNotNil(t, Detect("-/usr/local/bin/bash"))
	assertNotNil(t, Detect("-zsh"))
	assertNotNil(t, Detect("-/bin/zsh"))
	assertNotNil(t, Detect("-/usr/local/bin/zsh"))
}

func assertNotNil(t *testing.T, a Shell) {
	if a == nil {
		t.Error("Expected not to be nil")
	}
}

func assertEqual(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("Expected \"%v\" to equal \"%v\"", b, a)
	}
}
