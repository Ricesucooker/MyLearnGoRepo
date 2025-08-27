package main

import "testing"

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Top", "")
		want := "Hello, Top"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Top", "Japanese")
		want := "こんにちは, Top"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("Top", "Chinese")
		want := "你好, Top"

		assertCorrectMessage(t, got, want)
	})

}

//hello world fundamentals
//https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

// Key take away:
// Discipline
// Let's go over the cycle again
// Write a test
// Make the compiler pass
// Run the test, see that it fails and check the error message is meaningful
// Write enough code to make the test pass
// Refactor
