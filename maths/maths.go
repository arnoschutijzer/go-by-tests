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
	for _, slice := range slices {
		sums = append(sums, SumSlice(slice))
	}
	return sums
}

func SumAllTails(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, SumSlice(tail))
		}
	}
	return sums
}
