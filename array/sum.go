package array

// Takes an array of 5 numbers and returns their sum
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(slicesOfNumbers ...[]int) (returnSlice []int) {

	for _, sliceOfNumbers := range slicesOfNumbers {
		returnSlice = append(returnSlice, Sum(sliceOfNumbers))
	}
	return
}

func sumTails(slice []int) int {
	if len(slice) < 1 {
		return 0
	}
	return Sum(slice[1:])
}

func SumAllTails(slicesOfNumbers ...[]int) (returnSlice []int) {

	for _, sliceOfNumbers := range slicesOfNumbers {
		returnSlice = append(returnSlice, sumTails(sliceOfNumbers))
	}

	return
}
