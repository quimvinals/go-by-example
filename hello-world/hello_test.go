package main

import "testing"

func assertCorrectMessage(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q but expected %q", got, expected)
	}
}

func TestSayHello(t *testing.T) {
	t.Run("saying hello to custom people", func(t *testing.T) {
		got := sayHello("Quim")
		expected := "Hello, Quim"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("say 'Hello, World' when no name provided", func(t *testing.T) {
		got := sayHello("")
		expected := "Hello, World"

		assertCorrectMessage(t, got, expected)
	})
}
