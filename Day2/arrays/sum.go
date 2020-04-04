package arrays

func Sum(numbers [5]int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
