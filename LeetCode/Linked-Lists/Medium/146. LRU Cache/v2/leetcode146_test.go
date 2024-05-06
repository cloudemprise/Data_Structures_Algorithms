package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	// Test case 1: Get non-existent key
	if cache.Get(1) != -1 {
		t.Errorf("Test case 1 failed: expected -1, got %d", cache.Get(1))
	}

	// Test case 2: Put new key-value pair
	cache.Put(1, 1)
	if cache.Get(1) != 1 {
		t.Errorf("Test case 2 failed: expected 1, got %d", cache.Get(1))
	}

	// Test case 3: Put new key-value pair and evict least recently used
	cache.Put(2, 2)
	cache.Put(3, 3)
	if cache.Get(1) != -1 {
		t.Errorf("Test case 3 failed: expected -1, got %d", cache.Get(1))
	}
	if cache.Get(2) != 2 {
		t.Errorf("Test case 3 failed: expected 2, got %d", cache.Get(2))
	}
	if cache.Get(3) != 3 {
		t.Errorf("Test case 3 failed: expected 3, got %d", cache.Get(3))
	}

	// Test case 4: Update existing key-value pair
	cache.Put(2, 4)
	if cache.Get(2) != 4 {
		t.Errorf("Test case 4 failed: expected 4, got %d", cache.Get(2))
	}
}
