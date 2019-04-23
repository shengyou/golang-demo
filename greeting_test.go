package main

import "testing"

func TestGroupingYourTest(t *testing.T) {

	assertEqual := func(t *testing.T, want, got string) {
		t.Helper()
		if got != want {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	}

	t.Run("Print 'Hello!'", func(t *testing.T) {
		want := "Hello!"
		got := hello()

		assertEqual(t, want, got)
	})

	t.Run("Greeting to people", func(t *testing.T) {
		want := "Greeting, Shengyou"
		got := greeting("Shengyou", "")

		assertEqual(t, want, got)
	})

	t.Run("Greeting with empty parameter", func(t *testing.T) {
		want := "Greeting, everyone"
		got := greeting("", "")

		assertEqual(t, want, got)
	})

	t.Run("Add support for Mandarin", func(t *testing.T) {
		want := "您好，Shengyou"
		got := greeting("Shengyou", "Mandarin")

		assertEqual(t, want, got)
	})

	t.Run("Add support for Hawaii", func(t *testing.T) {
		want := "Hola, Tom"
		got := greeting("Tom", "Hawaii")

		assertEqual(t, want, got)
	})
}
