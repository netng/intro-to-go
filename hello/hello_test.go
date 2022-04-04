package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}

	t.Run("say hello 'Hello, world' when empty string provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("jono", "")
		want := "Hello, jono"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("jono", "Spanish")
		want := "Hola, jono"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("jono", "French")
		want := "Bonjour, jono"

		assertCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	got := Hello("jono", "French")
	fmt.Println(got)
	// Output: Bonjour, jono
}
