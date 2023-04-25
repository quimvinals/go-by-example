package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(arrays ...[]int) []int {
	var sums []int

	for _, array := range arrays {
		sums = append(sums, Sum(array))
	}

	return sums
}

func SumAllTails(arrays ...[]int) []int {
	var sums []int

	for _, array := range arrays {
		tail := array[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}