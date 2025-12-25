package algorithms

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "unsorted", input: []int{5, 2, 4, 3, 1}, want: []int{1, 2, 3, 4, 5}},
		{name: "already sorted", input: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "with duplicates", input: []int{4, 2, 2, 1}, want: []int{1, 2, 2, 4}},
		{name: "with negatives", input: []int{0, -3, 2, -1}, want: []int{-3, -1, 0, 2}},
		{name: "empty", input: []int{}, want: []int{}},
		{name: "nil", input: nil, want: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := append([]int(nil), tc.input...)
			got := BubbleSort(tc.input)

			if !slices.Equal(got, tc.want) {
				t.Fatalf("BubbleSort(%v) = %v, want %v", tc.input, got, tc.want)
			}

			if !slices.Equal(tc.input, original) {
				t.Fatalf("BubbleSort modified input: got %v, want %v", tc.input, original)
			}
		})
	}
}
