package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	var sums []int

	for _, v := range arrays {
		sums = append(sums, Sum(v))
	}
	return sums
}

func SumTails(arrays ...[]int) []int {
	var sums []int

	for _, v := range arrays {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}