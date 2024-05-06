package Heap

import (
	"reflect"
	"testing"
)

func TestHeap_heapifyUp(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		index    int
		expected Heap
	}{
		{
			name:     "test1",
			input:    []int{2, 3, 4, 5, 6, 7, 1},
			index:    6,
			expected: Heap{[]int{1, 3, 2, 5, 6, 7, 4}},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			h := &Heap{items: tt.input}
			h.heapifyUp(tt.index)

			if !reflect.DeepEqual(h, &tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, h)
			}
		})
	}
}

func TestHeap_Insert(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected Heap
	}{
		{
			name:     "test1",
			input:    []int{7, 6, 5, 4, 3, 2, 1},
			expected: Heap{[]int{1, 4, 2, 7, 5, 6, 3}},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			h := &Heap{}
			for _, v := range tt.input {
				h.Insert(v)
			}

			if !reflect.DeepEqual(h, &tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, h)
			}
		})
	}
}

func TestHeap_heapifyDown(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		index    int
		expected Heap
	}{
		{
			name:     "test1",
			input:    []int{3, 4, 2, 7, 5, 6},
			index:    0,
			expected: Heap{[]int{2, 4, 3, 7, 5, 6}},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			h := &Heap{items: tt.input}
			h.heapifyDown(tt.index)

			if !reflect.DeepEqual(h, &tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, h)
			}
		})
	}
}

func TestHeap_Pop(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected Heap
	}{
		{
			name:     "test1",
			input:    []int{1, 4, 2, 7, 5, 6, 3},
			expected: Heap{[]int{2, 4, 3, 7, 5, 6}},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			h := &Heap{items: tt.input}
			h.Pop()

			if !reflect.DeepEqual(h, &tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, h)
			}
		})
	}
}
