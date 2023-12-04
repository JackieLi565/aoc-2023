package utils

// sum of ints in an array
func Sum(arr *[]int) int {
	total := 0

	for _, num := range *arr {
		total += num
	}

	return total
}