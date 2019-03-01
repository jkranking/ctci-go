package ch01

import "testing"

func TestIsUnique(t *testing.T) {
	if IsUnique("unique") {
		t.Error("'unique' is not a unique string")
	}

	if !IsUnique("uniqke") {
		t.Error("'uniqke' is a unique string")
	}
}
