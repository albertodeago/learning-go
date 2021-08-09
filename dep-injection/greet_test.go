package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Johnny")

	got := buffer.String()
	want := "Hello, Johnny"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
