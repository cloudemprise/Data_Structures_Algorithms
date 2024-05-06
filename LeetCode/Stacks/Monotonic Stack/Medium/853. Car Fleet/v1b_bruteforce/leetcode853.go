package main

import (
	"fmt"
	"sort"
)

// Time Complexity : O(nlogn)
// Space Complexity: O(n)

// State of each car.
type car struct {
	distance int     // [miles] to final destination.
	duration float64 // [hours] until arrival.
}

// Stringer Interface
func (c car) String() string {
	return fmt.Sprintf("(%d/%.1f)", c.distance, c.duration)
}

// carFleet calculates how many individual fleets (groups of cars) arrive at the destination.
func carFleet(target int, position []int, speed []int) int {
	ln := len(position)

	// STEP 1:
	// We need to construct an ordered list of cars in terms of their distance
	// to the final destination and indicate their estimated time of arrival.

	// Build a list of the state of each car.
	cars := make([]car, ln)
	for i := 0; i < ln; i++ {
		cars[i] = car{
			distance: target - position[i],
			duration: float64(target-position[i]) / float64(speed[i]),
		}
	}

	// Sort the states in terms of their distance to the final destination, i.e.
	// increasing order where cars[0] has the shortest distance to travel, etc.
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].distance < cars[j].distance
	})
	fmt.Println()
	fmt.Println("Sorted car distance/duration to the final destination:")
	fmt.Println("\n\t\t", cars)
	fmt.Println()

	// STEP 2:
	// Determine if each car requires more time ('duration') than the one ahead
	// of it; if so, then there is no catch-up between them and the total fleet
	// number (groups of arriving  cars) increases by one. To look at it from a
	// different angle: if there are any cars behind another that
	// (theoretically) arrives before the car in front then that would imply
	// "hold-up" occurs and the total number of arriving groups of cars would
	// decrease thus reducing the total fleet number by one, where:
	// max(fleet) = total number of cars.

	// Use a standard stack to count the number of arriving groups of cars.
	stack := make([]car, 0)
	// Determine if the car in front is unhindered, i.e. not held up.
	for i := 0; i < ln; i++ {

		// If so, push it onto the stack.
		if (i == 0) || cars[i].duration > stack[len(stack)-1].duration {
			stack = append(stack, cars[i])
		}
	}
	// The stack holds only those cars in-front of each fleet.
	return len(stack)

}

///

func main() {

	var position []int
	var speed []int
	var target int

	var want int
	var got int

	///

	fmt.Println("TEST1:")
	target = 12
	fmt.Printf("target \t\t %#v \t\t\t [miles]\n", target)
	position = []int{10, 8, 0, 5, 3}
	fmt.Printf("position \t %2v \t [miles]\n", position)
	speed = []int{2, 4, 1, 1, 3}
	fmt.Printf("speed \t\t %2v \t [miles/hour]\n", speed)

	want = 3
	got = carFleet(target, position, speed)
	fmt.Printf("Want: %v \t Got: %v\n", want, got)
	fmt.Println("===============================")
	fmt.Println()

	///

	fmt.Println("TEST2:")
	target = 100
	fmt.Printf("target \t\t %#v \t\t [miles]\n", target)
	position = []int{0, 2, 4}
	fmt.Printf("position \t %2v \t [miles]\n", position)
	speed = []int{4, 2, 1}
	fmt.Printf("speed \t\t %2v \t [miles/hour]\n", speed)

	want = 1
	got = carFleet(target, position, speed)
	fmt.Printf("Want: %v \t Got: %v\n", want, got)
	fmt.Println("===============================")
	fmt.Println()

	///

	fmt.Println("TEST3:")
	target = 10
	fmt.Printf("target \t\t %#v \t\t [miles]\n", target)
	position = []int{6, 8}
	fmt.Printf("position \t %2v \t [miles]\n", position)
	speed = []int{3, 2}
	fmt.Printf("speed \t\t %2v \t [miles/hour]\n", speed)

	want = 2
	got = carFleet(target, position, speed)
	fmt.Printf("Want: %v \t Got: %v\n", want, got)
	fmt.Println("===============================")
	fmt.Println()

	///

	fmt.Println("TEST4:")
	target = 10
	fmt.Printf("target \t\t %#v \t\t [miles]\n", target)
	position = []int{0, 4, 2}
	fmt.Printf("position \t %2v \t [miles]\n", position)
	speed = []int{2, 1, 3}
	fmt.Printf("speed \t\t %2v \t [miles/hour]\n", speed)

	want = 1
	got = carFleet(target, position, speed)
	fmt.Printf("Want: %v \t Got: %v\n", want, got)
	fmt.Println("===============================")
	fmt.Println()

	///

	fmt.Println("TEST5:")
	target = 20
	fmt.Printf("target \t\t %#v \t\t [miles]\n", target)
	position = []int{0, 10, 15, 5}
	fmt.Printf("position \t %2v \t [miles]\n", position)
	speed = []int{5, 2, 1, 3}
	fmt.Printf("speed \t\t %2v \t [miles/hour]\n", speed)

	want = 1
	got = carFleet(target, position, speed)
	fmt.Printf("Want: %v \t Got: %v\n", want, got)
	fmt.Println("===============================")
	fmt.Println()
}
