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
	head := l.Head

	for head != nil {
		fmt.Fprint(w, head.Value, "-->")
		head = head.Next
	}
}
func (l *LinkedList) Search(target int) *listnode.ListNode {
	head := l.Head

	for head != nil {
		if head.Value == target {
			return head
		}
		head = head.Next
	}

	return nil
}
