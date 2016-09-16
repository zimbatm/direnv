package env

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	diff := &Diff{map[string]string{"FOO": "bar"}, map[string]string{"BAR": "baz"}}

	out := diff.Serialize()

	diff2, err := LoadDiff(out)
	if err != nil {
		t.Error("parse error", err)
	}

	if len(diff2.Prev) != 1 {
		t.Error("len(diff2.prev) != 1", len(diff2.Prev))
	}

	if len(diff2.Next) != 1 {
		t.Error("len(diff2.next) != 0", len(diff2.Next))
	}
}

// Issue #114
// Check that empty environment variables correctly appear in the diff
func TestDiffEmptyValue(t *testing.T) {
	before := Env{}
	after := Env{"FOO": ""}

	diff := BuildDiff(before, after)

	if !reflect.DeepEqual(diff.Next, map[string]string(after)) {
		t.Errorf("diff.Next != after (%#+v != %#+v)", diff.Next, after)
	}
}

func TestIgnoredEnv(t *testing.T) {
	if !ignoredEnv(DIRENV_BASH) {
		t.Fail()
	}
	if ignoredEnv(DIRENV_DIFF) {
		t.Fail()
	}
	if !ignoredEnv("_") {
		t.Fail()
	}
	if !ignoredEnv("__fish_foo") {
		t.Fail()
	}
	if !ignoredEnv("__fishx") {
		t.Fail()
	}
}
