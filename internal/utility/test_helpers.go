package utility

import (
	"slices"
	"testing"
)

func AssertEquality[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v, want %v", got, want)
	}
}

func AssertSliceEquality[E comparable, T ~[]E](t testing.TB, got, want T) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want %v", got, want)
	}
}
