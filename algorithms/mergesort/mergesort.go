package mergesort

func MergeSort(input []int) []int {
	length := len(input)

	if length <= 1 {
		return input
	}

	leftSide := MergeSort(input[:length/2])
	rightSide := MergeSort(input[length/2:])

	output := make([]int, length)

	var leftIndex, rightIndex, outputIndex int

	for leftIndex < len(leftSide) {
		for rightIndex < len(rightSide) && rightSide[rightIndex] <= leftSide[leftIndex] {
			output[outputIndex] = rightSide[rightIndex]
			outputIndex++
			rightIndex++
		}
		output[outputIndex] = leftSide[leftIndex]
		outputIndex++
		leftIndex++
	}

	for rightIndex < len(rightSide) {
		output[outputIndex] = rightSide[rightIndex]
		outputIndex++
		rightIndex++
	}

	return output
}
