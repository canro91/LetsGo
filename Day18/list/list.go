package list

import (
	"errors"
)

var (
	EmptyListError = errors.New("Empty list")
)

type Node struct {
	Value int
	next  *Node
}

type List struct {
	head *Node
}

func (l *List) AddFirst(node *Node) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *List) AddLast(node *Node) {
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

func (l *List) RemoveFirst() error {
	if l.head == nil {
		return EmptyListError
	}

	l.head = l.head.next
	return nil
}

func (l *List) RemoveLast() error {
	if l.head == nil {
		return EmptyListError
	}
	if l.head.next == nil {
		l.head = nil
		return nil
	}

	var previous *Node
	current := l.head
	for current.next != nil {
		previous = current
		current = current.next
	}
	previous.next = nil
	return nil
}
