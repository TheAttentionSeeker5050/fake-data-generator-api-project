package main

import (
	"testing"
)

func TestFooBar(t *testing.T) {

	// name variables
	var foo string = "foo"
	var bar string = "bar"

	// test if foo is equal to foo and bar is equal to bar
	if "foo" != foo {
		t.Errorf("foo is not equal to %s", foo)
	}

	if "bar" != bar {
		t.Errorf("bar is not equal to %s", bar)
	}
}
