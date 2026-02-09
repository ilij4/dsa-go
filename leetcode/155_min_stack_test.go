package leetcode

import "testing"

func TestMinStackSequence(t *testing.T) {
	s := Constructor()

	s.Push(-2)
	s.Push(0)
	s.Push(-3)

	if got := s.GetMin(); got != -3 {
		t.Fatalf("GetMin() after pushes = %d, want -3", got)
	}

	s.Pop()

	if got := s.Top(); got != 0 {
		t.Fatalf("Top() after pop = %d, want 0", got)
	}

	if got := s.GetMin(); got != -2 {
		t.Fatalf("GetMin() after pop = %d, want -2", got)
	}
}
