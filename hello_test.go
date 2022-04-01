package main

import "testing"

// requirements:
// # when we run hello.go
// # it should return string "Hello, ${name}"

func TestHello(t *testing.T) {
	got := Hello("asep")
	want := "Hello, asep"

	if got != want {
		t.Errorf("\n got: %q\n want: %q", got, want)
	}
}

