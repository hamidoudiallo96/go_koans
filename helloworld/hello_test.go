package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Hamidou", "")
		want := "Hello, Hamidou"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when no name is provided", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Shakira", "Spanish")
		want := "Hola, Shakira"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Antoine", "French")
		want := "Bonjour, Antoine"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Hisoka", "Japanese")
		want := "Konichiwa, Hisoka"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
