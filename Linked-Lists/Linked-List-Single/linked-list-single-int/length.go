package list

// Length traverses a list forward and returns the node count.
func (l *List) Length() int {
	var length int
	current := l.head
	for current != nil {
		current = current.next
		length++
	}
	return length
}
