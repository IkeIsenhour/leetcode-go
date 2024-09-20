package linkedlist

import (
	"fmt"
	"io"

	"github.com/IkeIsenhour/leetcode-go/internal/structure/listnode"
)

type LinkedList struct {
	Head *listnode.ListNode
}

func NewLinkedList(head *listnode.ListNode) *LinkedList {
	return &LinkedList{
		Head: head,
	}
}

func (l *LinkedList) Print(w io.Writer) {
	temp := l.Head

	for temp != nil {
		fmt.Fprint(w, temp.Value, "-->")
		temp = temp.Next
	}
}
