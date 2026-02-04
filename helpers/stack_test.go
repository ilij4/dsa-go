package helpers

import "testing"

func TestStackPopMutates(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)

	v, ok := s.Pop()
	if !ok {
		t.Fatalf("Pop() ok = false, want true")
	}
	if v != 2 {
		t.Fatalf("Pop() value = %d, want 2", v)
	}
	if got := s.Len(); got != 1 {
		t.Fatalf("Len() after Pop() = %d, want 1", got)
	}
}

func TestStackLIFOAndEmpty(t *testing.T) {
	s := NewStack[string]()
	s.Push("a")
	s.Push("b")

	v, ok := s.Pop()
	if !ok || v != "b" {
		t.Fatalf("Pop() = (%q, %v), want (%q, true)", v, ok, "b")
	}
	v, ok = s.Pop()
	if !ok || v != "a" {
		t.Fatalf("Pop() = (%q, %v), want (%q, true)", v, ok, "a")
	}
	if !s.IsEmpty() {
		t.Fatalf("IsEmpty() = false, want true")
	}

	_, ok = s.Pop()
	if ok {
		t.Fatalf("Pop() on empty ok = true, want false")
	}
}
