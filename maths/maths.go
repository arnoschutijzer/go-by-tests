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

// A variadic function!
func SumAll(slices ...[]int) (sums []int) {
	var totals []int
	for _, slice := range slices {
		totals = append(totals, SumSlice(slice))
	}

	return totals
}
