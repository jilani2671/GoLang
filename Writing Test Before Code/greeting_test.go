package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	t.Run("Say, string is filled with an argument", func(t *testing.T) {
		got := Greetings("Abrar")
		want := "Hello,Abrar"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Say, 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Greetings("")
		want := "Hello,World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
