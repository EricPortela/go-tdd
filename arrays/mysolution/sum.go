package mysolution

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(sliceOfNumbers ...[]int) []int {
	sum := make([]int, 0)

	for _, numbers := range sliceOfNumbers {
		sum = append(sum, Sum(numbers))
	}

	return sum
}

func SumAllTails(sliceOfNumbers ...[]int) []int {
	var sumTails []int

	for _, numbers := range sliceOfNumbers {

		if len(numbers) > 0 {
			sumTails = append(sumTails, Sum(numbers[1:]))
		} else {
			sumTails = append(sumTails, 0)
		}
	}

	return sumTails
}
