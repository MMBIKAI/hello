package main

import "testing"

func TestHi(t *testing.T) {

	got := hi("Chris")
	want := "Hello, Chris"

	if got != want {

		t.Errorf("got %q want %q", got, want)
	}
}
