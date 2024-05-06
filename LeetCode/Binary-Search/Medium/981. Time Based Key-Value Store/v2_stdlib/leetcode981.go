package main

import (
	"cmp"
	"fmt"
	"slices"
)

/*
The crux of this problem is the demonstration of the taking note of a most recent
solution found so far within the binary search algorithm as demonstrated by
Get() method.
*/

// Event value and its timestamp.
type record struct {
	value     string
	timestamp int
}

// Map a list of records.
type TimeMap struct {
	data map[string][]record
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]record),
	}
}

// Set stores a record under key within a list.
func (tm *TimeMap) Set(key string, value string, timestamp int) {

	// If the key doesn't yet exists.
	if _, ok := tm.data[key]; !ok {
		tm.data[key] = []record{} // allocate an empty list
	}
	// Append new record to key.
	tm.data[key] = append(tm.data[key], record{
		value:     value,
		timestamp: timestamp,
	})
}

// Get retrieves a most-recent previous event record.
func (tm *TimeMap) Get(key string, timestamp int) string {

	// Check if records exist.
	if records, ok := tm.data[key]; ok {

		// The most-recent previous value.
		var result string

		// Perform a binary search on the records list.
		target := record{timestamp: timestamp}
		index, isFound := slices.BinarySearchFunc(records, target, func(a, b record) int {
			return cmp.Compare(a.timestamp, b.timestamp)
		})

		// Check if the record exists.
		if isFound {
			result = records[index].value
		} else if index > 0 {
			// Or most-recent previous record exists.
			result = records[index-1].value
		}

		return result
	}
	return "" // key doesn't exist.
}

///

func main() {

	cache := Constructor()

	cache.Set("love", "high", 10)
	cache.Set("love", "low", 20)

	fmt.Printf("Get(\"love\",  5) \t Want: \"\" \t Got: %#v \n\n", cache.Get("love", 5))
	fmt.Printf("Get(\"love\", 10) \t Want: high \t Got: %#v \n\n", cache.Get("love", 10))
	fmt.Printf("Get(\"love\", 15) \t Want: high \t Got: %#v \n\n", cache.Get("love", 15))
	fmt.Printf("Get(\"love\", 20) \t Want: low \t Got: %#v \n\n", cache.Get("love", 20))
	fmt.Printf("Get(\"love\", 25) \t Want: low \t Got: %#v \n\n", cache.Get("love", 25))

}
