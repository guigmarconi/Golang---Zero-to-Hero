package main

func Sum(numbers []int) int {

	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func SumAll(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSumTails ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSumTails {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
