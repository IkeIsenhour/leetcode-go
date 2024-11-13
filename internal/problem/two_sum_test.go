package problem

import (
	"testing"

	"github.com/IkeIsenhour/leetcode-go/internal/utility"
)

func TestTwoSum(t *testing.T) {
	arr := []int{2, 21, 303, 7}

	t.Run("successful TwoSum #1", func(t *testing.T) {
		got := TwoSum(arr, 9)
		want := []int{0, 3}

		utility.AssertSliceEquality(t, got, want)
	})
}
