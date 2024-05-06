package queue

/*
A FIFO queue ADT using a slice as the backing data structure. No type
assertion idiom is necessarily needed at this level.
*/

type Queue struct {
	items []any
}

//*-*

func New() *Queue {
	return &Queue{}
}

//*-*

func (q *Queue) EnQueue(item any) {
	q.items = append(q.items, item)
}

//*-*

/* // Just another way of doing it.
func (q *Queue) DeQueue() any {
	if ln := len(q.items); ln > 0 {
		tmp := q.items[0]
		q.items = q.items[1:]
		return tmp
	}
	log.Fatal("DeQueue() : empty queue")
	return nil
} */

func (q *Queue) DeQueue() any {
	if ln := len(q.items); ln > 0 {
		defer func() {
			q.items = q.items[1:]
		}()
		return q.items[0]
	}
	panic("DeQueue() : empty queue")
}

//*-*

func (q *Queue) Size() int {
	return len(q.items)
}

//*-*

func (q *Queue) Front() any {
	return q.items[0]
}

//*-*

func (q *Queue) IsEmpty() bool {
	return !(len(q.items) > 0)
}

//*-*

func (q *Queue) Reset() {
	//q.items = q.items[:0]
	q.items = nil
}

//*-*
