package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jericho")
	want := "Hello, Jericho"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
