package main

import "fmt"

func findJudge(n int, trust [][]int) int {

	// A person can trust more than one person but a Judge trusts no one,
	// including himself, i.e. a Judge has ZERO trust relations.

	// Count how many times a person isTrusted.
	// Count how many times a person hasTrusted.
	// Start at index 1 to coincide with each person.
	isTrusted := make([]int, n+1)
	hasTrusted := make([]int, n+1)
	for _, trustRelation := range trust {

		truster := trustRelation[0] // person who has trust.
		trustee := trustRelation[1] // person who is trusted.

		// Count the amount of times a person is trusted. A possible Judge
		// will have 'n-1' trust-relations(votes).
		isTrusted[trustee]++

		// Count the amount of times a person has trusted another. A possible
		// Judge has no trust.
		hasTrusted[truster]++
	}

	// Search for a possible Judge.
	requiredTrust := n - 1
	for person := 1; person <= n; person++ {

		if hasTrusted[person] == 0 && isTrusted[person] == requiredTrust {
			judge := person
			return judge
		}
	}
	return -1 // no Judge.
}

func main() {

	var numPeople int
	var trustRelations [][]int

	numPeople = 2
	trustRelations = [][]int{
		{1, 2},
	}
	fmt.Println("There are ", numPeople, "people in the town.")
	fmt.Printf("These are the trust relations: %v \n", trustRelations)
	fmt.Printf("Want: 2 \t Got: %v \n", findJudge(numPeople, trustRelations))
	fmt.Println()

	///

	numPeople = 3
	trustRelations = [][]int{
		{1, 3},
		{2, 3},
	}
	fmt.Println("There are ", numPeople, "people in the town.")
	fmt.Printf("These are the trust relations: %v \n", trustRelations)
	fmt.Printf("Want: 3 \t Got: %v \n", findJudge(numPeople, trustRelations))
	fmt.Println()

	///

	numPeople = 3
	trustRelations = [][]int{
		{1, 3},
		{2, 3},
		{3, 1},
	}
	fmt.Println("There are ", numPeople, "people in the town.")
	fmt.Printf("These are the trust relations: %v \n", trustRelations)
	fmt.Printf("Want: -1 \t Got: %v \n", findJudge(numPeople, trustRelations))
	fmt.Println()

	///

	numPeople = 4
	trustRelations = [][]int{
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{4, 3},
	}
	// [[1,3],[1,4],[2,3],[2,4],[4,3]]
	fmt.Println("There are ", numPeople, "people in the town.")
	fmt.Printf("These are the trust relations: %v \n", trustRelations)
	fmt.Printf("Want: 3 \t Got: %v \n", findJudge(numPeople, trustRelations))
	fmt.Println()

	///

}
