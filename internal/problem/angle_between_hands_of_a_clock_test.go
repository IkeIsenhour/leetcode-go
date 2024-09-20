package problem

import (
	"testing"

	"github.com/IkeIsenhour/leetcode-go/internal/utility"
)

func TestAngleClock(t *testing.T) {

	t.Run("check the angle of 3:15", func(t *testing.T) {
		got := AngleClock(3, 15)
		want := 7.50

		utility.AssertEquality(t, got, want)
	})

	t.Run("check the angle when hands are in the same location", func(t *testing.T) {
		got := AngleClock(12, 0)
		want := 0.0

		utility.AssertEquality(t, got, want)
	})
}
