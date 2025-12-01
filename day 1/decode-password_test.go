package main

import "testing"

func TestDecodePassword(t *testing.T) {
	t.Run("Run example given", func(t *testing.T) {
		got := decodePassword("example.txt")
		want := 3
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
