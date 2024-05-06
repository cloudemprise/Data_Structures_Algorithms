package queue

// A FIFO queue ADT using a slice as the backing data structure.

// Queue definition slice
type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{}
}

///

func (q *Queue) EnQueue(item int) {
	q.items = append(q.items, item) // HEAD @ index 0
}

func (q *Queue) DeQueue() int {
	tmp := q.items[0]
	q.items = q.items[1:]
	return tmp
}

/* func (q *Queue) DeQueue() int {
	defer func() {
		q.items = q.items[1:]
	}()
	return q.items[0]
} */

///

func (q *Queue) Peek() int {
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	//return !(len(q.items) > 0)
	return len(q.items) == 0
}

func (q *Queue) Length() int {
	return len(q.items)
}

func (q *Queue) Reset() {
	//q.items = q.items[:0]
	q.items = nil
}

///

/* func (q *Queue) String() string {
	var qStr strings.Builder
	qStr.WriteString("[")
	for _, item := range q.items {
		fmt.Fprintf(&qStr, " %v", item)
	}
	qStr.WriteString(" ]")
	return qStr.String()
} */

///

/* func main() {

	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	queue := NewQueue()
	fmt.Println("An empty queue: ")
	fmt.Println(queue)

	fmt.Println("Enqueuing numbers:")
	for _, num := range nums {
		queue.EnQueue(num)
		fmt.Println(queue)
	}

	fmt.Println("Dequeuing numbers:")
	for i := 5; i > 0; i-- {
		fmt.Printf("\ndeQueue %v ", queue.DeQueue())
		fmt.Println(queue)
	}

	fmt.Printf("Size of Queue = %v\n", queue.Length())
	queue.Reset()
	fmt.Printf("Size of Queue = %v\n", queue.Length())
	fmt.Println(queue)
}
*/
