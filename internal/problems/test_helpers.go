package problems

import "testing"

func assertEquality[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v, want %v", got, want)
	}
}
