package main

/*
1. say 'Hello, World'
2. say 'Hello, ${name}'
3. say 'Hello, World' when empty string supplied as argument
*/

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("nan")
		want := "Hello, nan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when empty string supplied as argument", func(t *testing.T) {
		got := Hello("nan")
		want := "Hello, nan"
		assertCorrectMessage(t, got, want)
	})
}
