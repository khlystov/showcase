package session

import (
	"testing"
)

func TestSet(t *testing.T) {
	Store.Set("test")
	if !Store.Check("test") {
		t.Error()
	}

	if Store.Check("test_no") {
		t.Error()
	}
}

func TestCheck(t *testing.T) {
	Store.Set("check")
	if !Store.Check("check") {
		t.Error()
	}

	if Store.Check("test_no") {
		t.Error()
	}
}
