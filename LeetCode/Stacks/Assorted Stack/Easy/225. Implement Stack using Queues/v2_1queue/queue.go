package main

// A FIFO queue ADT using a slice as the backing data structure.

// queue definition slice
type queue struct {
	items []int
}

func newQueue() *queue {
	return &queue{}
}

///

func (q *queue) enQueue(item int) {
	q.items = append(q.items, item) // HEAD @ index 0
}

func (q *queue) deQueue() int {
	tmp := q.items[0]
	q.items = q.items[1:]
	return tmp
}

/* func (q *queue) deQueue() int {
	defer func() {
		q.items = q.items[1:]
	}()
	return q.items[0]
} */

///

func (q *queue) peek() int {
	return q.items[0]
}

func (q *queue) isEmpty() bool {
	//return !(len(q.items) > 0)
	return len(q.items) == 0
}

func (q *queue) length() int {
	return len(q.items)
}

func (q *queue) reset() {
	//q.items = q.items[:0]
	q.items = nil
}
