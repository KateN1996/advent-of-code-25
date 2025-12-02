package main

import "testing"

func TestDecodePassword(t *testing.T) {
	t.Run("Day 1 part 1: Run example given", func(t *testing.T) {
		got := decodePassword("example.txt")
		want := 3
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Day 1 part 2: run example given", func(t *testing.T) {

		got := decodePassword2("example.txt")
		want := 6
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
