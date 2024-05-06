package main

import (
	"fmt"
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

		// To remember the most-recent previous value so far.
		var result string

		// Perform a binary search on the records list.
		lo, hi := 0, len(records)-1
		var mid int
		for lo <= hi {

			mid = (hi + lo) / 2

			// Search for most recent previous timestamp.
			if records[mid].timestamp <= timestamp {
				// Take note of the search result so far...
				result = records[mid].value
				// ...but consider if there is a timestamp of
				// a more recent previous event.
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return result
	}
	return "" // no record found.
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
