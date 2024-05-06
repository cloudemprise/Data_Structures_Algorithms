package main

import "fmt"

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type LRUCache struct {
	cap   int
	store map[int]*node
	head  *node
	tail  *node
}

func Constructor(cap int) LRUCache {
	head, tail := &node{}, &node{}
	head.next, tail.prev = tail, head
	return LRUCache{
		cap:   cap,
		store: make(map[int]*node),
		head:  head,
		tail:  tail,
	}
}

///

// helper function to remove a node from a DLL.
func (c *LRUCache) remove(n *node) {
	// Delete key-store allocation.
	delete(c.store, n.key)

	// Block out current node pointers.
	n.prev.next = n.next
	n.next.prev = n.prev
	// Clean dangling pointers for GC.
	n.next, n.prev = nil, nil
}

// helper function to insert a node into the Head of a DLL.
func (c *LRUCache) insert(n *node) {
	// Create new key-store allocation.
	c.store[n.key] = n

	// Bind new node pointers.
	n.next = c.head.next
	n.prev = c.head
	// Bind previous leader node pointer.
	c.head.next.prev = n
	// Update Head.
	c.head.next = n

}

///

func (c *LRUCache) Get(key int) int {
	// Check if key exists
	if n, exists := c.store[key]; exists {
		c.remove(n)
		c.insert(n)
		return n.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {

	// Check if key exists
	if n, exists := c.store[key]; exists {
		c.remove(n)
	}

	// NOTE!!! Above & Below logic needs to be in this sequence since
	// we only need to evict on a new Put() and not on a refresh.

	// Cache FULL? Delete LRU node.
	if len(c.store) == c.cap {
		c.remove(c.tail.prev)
	}

	// Allocate a new item.
	newNode := &node{
		key:   key,
		value: value,
	}
	// Register new item.
	c.insert(newNode)
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
