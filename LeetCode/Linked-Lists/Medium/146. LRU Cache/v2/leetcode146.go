package main

import "fmt"

// The fact that Head & Tail are actual dummy nodes means that their pointers
// never have to be reallocated but the nodes themselves need to be reallocated
// every time an operation is performed on them.

type node struct {
	key, value int
	prev, next *node
}

type LRUCache struct {
	cap        int
	head, tail *node // Are actual nodes and not just pointers to nodes.
	store      map[int]*node
}

func newLRUCache(head, tail *node, cap int) LRUCache {
	return LRUCache{
		head:  head,
		tail:  tail,
		store: make(map[int]*node),
		cap:   cap,
	}
}

func newNode(key, val int) *node {
	return &node{
		key:   key,
		value: val,
	}
}

func Constructor(capacity int) LRUCache {
	// Are actual nodes and not just pointers to nodes.
	head, tail := newNode(0, 0), newNode(0, 0)

	head.next = tail
	tail.prev = head
	return newLRUCache(head, tail, capacity)
}

func (c *LRUCache) remove(node *node) {
	delete(c.store, node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) insert(node *node) {
	c.store[node.key] = node
	next := c.head.next
	c.head.next = node
	node.prev = c.head
	next.prev = node
	node.next = next
}

// Get also does a reallocation
func (c *LRUCache) Get(key int) int {
	if n, ok := c.store[key]; ok {
		c.remove(n)
		c.insert(n)
		return n.value
	}

	return -1
}

// Does a reallocation regardless of the contents of the key-store, i.e. Put()
// of an already existing key is reallocated a new node.
func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.store[key]; ok {
		c.remove(c.store[key])
	}

	if len(c.store) == c.cap {
		c.remove(c.tail.prev)
	}

	c.insert(newNode(key, value))
}

///

func main() {

	//fmt.Println("Hello World")

	obj := Constructor(1)
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
