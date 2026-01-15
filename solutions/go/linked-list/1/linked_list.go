package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value any
	prev  *Node
	next  *Node
}

type List struct {
	first *Node
	last  *Node
}

func NewList(elements ...any) *List {
	if len(elements) == 0 {
		return &List{}
	}

	n := &Node{Value: elements[0]}

	curr := n

	for _, v := range elements[1:] {
		curr.next = &Node{Value: v, prev: curr}
		curr = curr.next
	}

	return &List{n, curr}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	if l.first == nil {
		l.first = &Node{Value: v}
		l.last = l.first
		return
	}

	l.first.prev = &Node{Value: v, next: l.first}
	l.first = l.first.prev
}

func (l *List) Push(v any) {
	if l.first == nil {
		l.first = &Node{Value: v}
		l.last = l.first
		return
	}

	l.last.next = &Node{Value: v, prev: l.last}
	l.last = l.last.next
}

func (l *List) Shift() (any, error) {
	if l.first == nil {
		return nil, errors.New("")
	}

	v := l.first.Value

	// Only 1 node
	if l.first.next == nil {
		l.first, l.last = nil, nil
	} else {
		l.first = l.first.next
		l.first.prev = nil
	}

	return v, nil
}

func (l *List) Pop() (any, error) {
	if l.last == nil {
		return nil, errors.New("")
	}

	v := l.last.Value

	if l.first.next == nil {
		l.first, l.last = nil, nil
	} else {
		l.last = l.last.prev
		l.last.next = nil
	}

	return v, nil
}

func (l *List) Reverse() {
	if l.first == nil {
		return
	}

	l.first, l.last = l.last, l.first

	n := l.last
	for n != nil {
		n.next, n.prev = n.prev, n.next
		n = n.prev
	}
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
