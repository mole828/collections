package sortedlist

type Element[T any] struct {
	value T
	next  *Element[T]
	prev  *Element[T]
}

func (element Element[T]) Value() T {
	return element.value
}

func (e Element[T]) Next() *Element[T] {
	return e.next
}

func (e Element[T]) Prev() *Element[T] {
	return e.prev
}

type SortedList[T any] struct {
	head, tail *Element[T]
	len        int
	compare    func(T, T) bool
}

func New[T any](compare func(T, T) bool) *SortedList[T] {
	return &SortedList[T]{
		head: &Element[T]{},
		// tail:    &Element[T]{},
		len:     0,
		compare: compare,
	}
}

func (list *SortedList[T]) Head() *Element[T] {
	return list.head.next
}

func (list *SortedList[T]) Tail() *Element[T] {
	return list.tail.prev
}

func (list *SortedList[T]) Add(value T) {
	e := &Element[T]{
		value: value,
	}
	a := list.head
	b := a.next
	if b == nil {
		a.next = e
		return
	}

	a = b
	b = b.next

	if list.compare(value, a.value) {
		list.head.next = e
		e.next = a
		a.prev = e
		return
	}

	for b != nil {
		if list.compare(a.value, value) && list.compare(value, b.value) {
			break
		}
		a = b
		b = b.next
	}

	a.next = e
	if a != list.head {
		e.prev = a
	}
	e.next = b
	if b != nil {
		b.prev = e
	}

	list.len += 1
}

func (list *SortedList[T]) Remove(is func(T) bool) {
	a := list.head
	b := a.next

	for b != nil {
		if is(b.value) {
			c := b.next

			b.prev = nil
			b.next = nil

			a.next = c
			if c != nil {
				c.prev = nil
				if a != list.head {
					c.prev = a
				}
			}

			b = c
		} else {
			a = b
			b = b.next
		}
	}
}
