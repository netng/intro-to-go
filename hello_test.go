package main

import "testing"

/*
1. say 'Hello, World' -> this test should pass first
2. say 'Hello, ${name}' -> this test should pass second
3. say 'Hello, World' when empty string supplied as argument -> this test is last
*/

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("jono")
		want := "Hello, jono"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' when empty string provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}
