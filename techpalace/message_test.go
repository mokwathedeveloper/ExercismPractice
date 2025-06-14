package main

import "testing"

func TestWelcomeMessage(t *testing.T) {
	got := WelcomeMessage("Jane")
	want := "Welcome to the Tech Palace, JANE"
	if got != want {
		t.Errorf("WelcomeMessage = %q, want %q", got, want)
	}
}

func TestAddBorder(t *testing.T) {
	got := AddBorder("Hello", 5)
	want := "*****\nHello\n*****"
	if got != want {
		t.Errorf("AddBorder = %q, want %q", got, want)
	}
}

func TestCleanUpMessage(t *testing.T) {
	input := `
***************
*  Hello Go!  *
***************
`
	got := CleanUpMessage(input)
	want := "Hello Go!"
	if got != want {
		t.Errorf("CleanUpMessage = %q, want %q", got, want)
	}
}
