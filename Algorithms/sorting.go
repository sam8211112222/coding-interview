package Algorithms

func BubbleSort(input []int) []int {

	reloop := true
	for reloop {
		reloop = false
		for i := 0; i < len(input)-1; i++ {
			if input[i] > input[i+1] {
				temp := input[i+1]
				input[i+1] = input[i]
				input[i] = temp
				reloop = true
			}
		}
	}
	return input
}

func SelectionSort(input []int) []int {

	for i := 0; i < len(input); i++ {
		minIndex := i
		for j := i + 1; j < len(input); j++ {
			if input[minIndex] > input[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			temp := input[i]
			input[i] = input[minIndex]
			input[minIndex] = temp
		}
	}

	return input
}
