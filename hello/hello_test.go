package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jon", "English")
		want := "Hello, Jon!"

		assertMessage(t, got, want)
	})

	t.Run("saying 'Hello, world!' when no name is provided", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world!"

		assertMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "¡Hola, Juan!"

		assertMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour, Jean!"

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
