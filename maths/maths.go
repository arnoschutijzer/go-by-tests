package maths

func Sum(elements [4]int) int {
	var sum int
	for _, number := range elements {
		sum += number
	}
	return sum
}

func SumSlice(elements []int) int {
	var sum int
	for _, number := range elements {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) (sums []int) {
	totals := make([]int, len(slices))

	for index, slice := range slices {
		totals[index] = SumSlice(slice)
	}

	return totals
}
