package main

// Sum all the numbers in a slice and return it
func Sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

// Sum all the numbers in a slice of slices, then return a slice with all the results
func SumAll(numbersToSum ...[]int) (sum []int) {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// Sum all the numbers in a slice of slices (except the first of each slice), then return a slice with all the results
func SumAllTails(numbersToSum ...[]int) (sum []int) {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}

	return sums
}
