package main

import "fmt"

func findJudge(n int, trust [][]int) int {

	// The thing to take note of is that everyone trusts(votes for) the Judge
	// but the Judge trusts(votes) for no one, including himself, i.e. the
	// Judge has NO trust relation.

	// Count all trust-relations per person.
	// Start at index 1 to coincide with each person.
	trustCount := make([]int, n+1)
	for _, trustRelation := range trust {

		truster := trustRelation[0] // person who has trust.
		trustee := trustRelation[1] // person who is trusted.

		// Increment all the trust-relations(votes). Therefore, a possible
		// Judge will have 'n-1' trust-relations(votes).
		trustCount[trustee]++

		// Decrement the trust count of those people who have a trust-relation.
		// Therefore, if a potential Judge has a trust-relation, i.e. trusts
		// someone, then this will exclude him from being a possible Judge,
		// i.e. will have less than 'n-1' trust-relations(votes).
		trustCount[truster]--
	}

	// Search for a possible Judge.
	requiredVotes := n - 1
	for person := 1; person <= n; person++ {

		if trustCount[person] == requiredVotes {
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
