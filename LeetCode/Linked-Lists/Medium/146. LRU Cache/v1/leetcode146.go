package main

import "fmt"

/*
This solution uses a doubly linked list to keep track of the order of the
elements in the cache, with the most recently used element at the front and the
least recently used element at the back. It also uses a map to provide O(1)
access to the elements in the cache.

The `Get` method first checks if the key is in the cache. If it is, it moves the
corresponding element to the front of the list and returns its value. If it is
not, it returns -1.

The `Put` method first checks if the key is already in the cache. If it is, it
updates the value of the corresponding element and moves it to the front of the
list. If it is not, it checks if the cache is at capacity. If it is, it removes
the least recently used element from the back of the list and from the cache. It
then adds the new element to the front of the list and to the cache.

The solution makes use of helper functions to insert and cut nodes from the
linked list.
*/

type list struct {
	head *node
	tail *node
}

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type LRUCache struct {
	cap   int
	list  *list
	store map[int]*node
}

func Constructor(cap int) LRUCache {
	return LRUCache{
		cap:   cap,
		list:  &list{},
		store: make(map[int]*node),
	}
}

///

// helper function to insert a node into the Head of a DLL.
func (l *list) insertHead(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = l.head
	} else {
		// Bind new node.
		n.next = l.head
		n.prev = nil // cautionary.
		// Bind old Head node.
		l.head.prev = n
		// Set Head pointer.
		l.head = n
	}
}

// helper function to cut out a node from a DLL.
func (l *list) cutNode(n *node) {

	switch {
	case l.tail == n && l.head == n:
		// Only one node.
		l.head, l.tail = nil, nil

	case l.tail == n && l.head != n:
		// More than one node.
		n.prev.next = nil // detach old Tail node.
		l.tail = n.prev   // update Tail pointer.

	default: // Middle of linked list.
		n.prev.next = n.next // block off current node.
		n.next.prev = n.prev
	}

	// Clean dangling pointers for GC.
	n.next, n.prev = nil, nil
}

///

func (c *LRUCache) Get(key int) int {
	// Check if key exists
	n, exists := c.store[key]
	if exists {
		// If node not MRU, make it so.
		if n != c.list.head {
			// Cut node out of list.
			c.list.cutNode(n)
			// Insert into Head position.
			c.list.insertHead(n)
		}
		return n.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// Check if key exists
	n, exists := c.store[key]

	if exists { // Update key-store:

		if n != c.list.head { // Update MRU:

			// _ = c.Get(key) // Lazy approach.

			// Cut node out of list.
			c.list.cutNode(n)
			// Insert into Head position.
			c.list.insertHead(n)
		}
		n.value = value // New updated value.

	} else { // New key:

		// Cache FULL? Delete LRU node.
		if len(c.store) == c.cap {
			// Delete LRU key-store item.
			delete(c.store, c.list.tail.key)
			// Discard Tail node.
			c.list.cutNode(c.list.tail)
		}

		// Allocate a new item.
		newNode := &node{
			key:   key,
			value: value,
		}

		// Register new key-store item.
		c.store[key] = newNode
		// Update MRU
		c.list.insertHead(newNode)
	}
}

///

func main() {

	//fmt.Println("Hello World")

	obj := Constructor(2)
	fmt.Println("Put(1, 1)")
	obj.Put(1, 1)
	fmt.Println("Put(2, 2)")
	obj.Put(2, 2)
	fmt.Printf("[Get(1)] Want: 1 \t Got: %v\n", obj.Get(1))
	fmt.Println("Put(3, 3)")
	obj.Put(3, 3)
	fmt.Printf("[Get(2)] Want: -1 \t Got: %v\n", obj.Get(2))
	fmt.Println("Put(4, 4)")
	obj.Put(4, 4)
	fmt.Printf("[Get(1)] Want: -1 \t Got: %v\n", obj.Get(1))
	fmt.Printf("[Get(3)] Want: 3 \t Got: %v\n", obj.Get(3))
	fmt.Printf("[Get(4)] Want: 4 \t Got: %v\n", obj.Get(4))

	//*-*

}

//*-*
