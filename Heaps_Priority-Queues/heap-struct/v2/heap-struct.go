package Heap

/*
Implementation of a Min-Heap with struct-slice backing ADT.
*/

type Heap struct {
	items []int
}

func (h *Heap) Insert(value int) {
	h.items = append(h.items, value)
	h.heapifyUp(len(h.items) - 1)
}

func (h *Heap) heapifyUp(index int) {
	parent := (index - 1) / 2
	for index > 0 && h.items[index] < h.items[parent] {
		h.items[index], h.items[parent] = h.items[parent], h.items[index]
		index = parent
		parent = (index - 1) / 2
	}
}

// heapifyDown recursively...
func (h *Heap) heapifyDown(index int) {
	left := 2*index + 1
	right := 2*index + 2
	current := index

	if left < len(h.items) && h.items[left] < h.items[current] {
		current = left
	}

	if right < len(h.items) && h.items[right] < h.items[current] {
		current = right
	}

	if current != index { // Base Case
		h.items[index], h.items[current] = h.items[current], h.items[index]
		h.heapifyDown(current) // Tail recursion
	}
}

func (h *Heap) Pop() int {
	topPriority := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.heapifyDown(0)
	return topPriority
}
