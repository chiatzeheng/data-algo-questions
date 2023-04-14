package main

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	n := len(people)
	sort.Ints(people)
	i, j := 0, n-1
	count := 0

	for i <= j {
		if people[i]+people[j] <= limit {
			i++
		}
		j--
		count++
	}

	return count
}
