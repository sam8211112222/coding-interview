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
