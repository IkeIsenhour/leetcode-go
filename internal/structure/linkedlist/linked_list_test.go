package linkedlist

import (
	"bytes"
	"testing"

	"github.com/IkeIsenhour/leetcode-go/internal/structure/listnode"
	"github.com/IkeIsenhour/leetcode-go/internal/utility"
)

func TestPrint(t *testing.T) {
	c := listnode.NewListNode(3, nil)
	b := listnode.NewListNode(2, c)
	a := listnode.NewListNode(1, b)
	ll := NewLinkedList(a)

	t.Run("print output", func(t *testing.T) {
		var output bytes.Buffer

		ll.Print(&output)
		got := output.String()
		want := "1-->2-->3-->"

		utility.AssertEquality(t, got, want)
	})
}

func TestSearch(t *testing.T) {
	c := listnode.NewListNode(3, nil)
	b := listnode.NewListNode(2, c)
	a := listnode.NewListNode(1, b)
	ll := NewLinkedList(a)

	t.Run("search success", func(t *testing.T) {
		node := ll.Search(3)

		got := node.Value
		want := 3

		utility.AssertEquality(t, got, want)
	})
}
