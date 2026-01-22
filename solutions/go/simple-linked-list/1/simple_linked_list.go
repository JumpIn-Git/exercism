package linkedlist

import "errors"

// Define the List and Element types here.
type List struct {
	first *Element
	last  *Element
	len   int
}
type Element struct {
	value int
	next  *Element
}

func New(elements []int) *List {
	if len(elements) == 0 {
		return &List{}
	}

	n := &Element{value: elements[0]}
	l := 1
	curr := n

	for _, v := range elements[1:] {
		curr.next = &Element{value: v}
		curr = curr.next
		l++
	}

	return &List{n, curr, l}
}

func (l *List) Size() int {
	return l.len
}

func (l *List) Push(element int) {
	if l.Size() == 0 {
		l.first = &Element{value: element}
		l.last = l.first
	} else {
		l.last.next = &Element{value: element}
		l.last = l.last.next
	}

	l.len++
}

func (l *List) Pop() (int, error) {
	if l.Size() == 0 {
		return 0, errors.New("Empty list")
	}

	if l.first.next == nil {
		v := l.first.value
		l.first, l.last = nil, nil
		l.len--
		return v, nil
	}

	var prev *Element
	curr := l.first

	for curr.next != nil {
		prev, curr = curr, curr.next
	}

	prev.next = nil
	l.last = prev
	l.len--

	return curr.value, nil
}

func (l *List) Array() []int {
	a := make([]int, 0, l.Size())
	if l.Size() == 0 {
		return a
	}

	n := l.first
	for n.next != nil {
		a = append(a, n.value)
		n = n.next
	}
	a = append(a, n.value)

	return a
}

func (l *List) Reverse() *List {
	if l.Size() <= 1 {
		return l
	}

	var prev *Element
	curr := l.first

	l.last = l.first

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.first = prev

	return l
}
