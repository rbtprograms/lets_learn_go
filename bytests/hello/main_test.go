package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrect := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bobby", "")
		want := "Hello, Bobby"
		assertCorrect(t, got, want)
	})

	t.Run("say Hello, World if no string is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrect(t, got, want)
	})

	t.Run("spanish support", func(t *testing.T) {
		got := Hello("Bobby", "Spanish")
		want := "Hola, Bobby"
		assertCorrect(t, got, want)
	})

	t.Run("french support", func(t *testing.T) {
		got := Hello("Bobby", "French")
		want := "Bonjour, Bobby"
		assertCorrect(t, got, want)
	})
}
