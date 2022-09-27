package main

import "testing"

func TestSayHello(t *testing.T) {
	got := sayHello("Quim")
	expected := "Hello, Quim"

	if got != expected {
		t.Errorf("got %q but expected %q", got, expected)
	}
}
