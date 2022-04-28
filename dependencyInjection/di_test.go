package main

import (
	"bytes"
	"testing"
)

func TestDi(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nan")

	got := buffer.String()
	want := "Hello, Nan"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
