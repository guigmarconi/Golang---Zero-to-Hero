package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Guigo", "")
		want := "Hello, Guigo"
		assertMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)

	})

	t.Run("in Portuguese", func(t *testing.T) {

		got := Hello("Mundo", "Portuguese")
		want := "Ola, Mundo"
		assertMessage(t, got, want)

	})

	t.Run("in French", func(t *testing.T) {

		got := Hello("Gustave", "French")
		want := "Bonjour, Gustave"
		assertMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {

		got := Hello("Miguel", "Spanish")
		want := "Hola, Miguel"
		assertMessage(t, got, want)

	})

}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
