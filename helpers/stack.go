package helpers

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() Stack[T] {
	items := make([]T, 0)

	return Stack[T]{items: items}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T

	n := len(s.items)
	if n == 0 {
		return zero, false
	}

	v := s.items[n-1]

	var z T
	s.items[n-1] = z
	s.items = s.items[:n-1]

	// Best, but overkil here
	// tmp := make([]T, length-1)
	// copy(tmp, s.items)
	// s.items = tmp

	return v, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T

	if s.IsEmpty() {
		return zero, false
	}

	return s.items[len(s.items)-1], true
}
