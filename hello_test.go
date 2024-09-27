package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	got := Greetings("Abrar")
	want := "Hello,Abrar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
