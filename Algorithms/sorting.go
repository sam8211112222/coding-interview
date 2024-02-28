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

// InsertionSort 實現插入排序算法
func InsertionSort(input []int) []int {
	// 從第二個元素開始遍歷切片
	for i := 1; i < len(input); i++ {
		temp := input[i]
		j := i - 1

		// 將當前元素與它前面的元素比較，如果前面的元素更大，則將它們向後移動
		for j >= 0 && input[j] > temp {
			input[j+1] = input[j]
			j--
		}
		// 將當前元素放到它應該在的位置
		input[j+1] = temp
	}
	return input
}

// MergeSort function that recursively sorts the input slice
func MergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	half := len(input) / 2
	left := MergeSort(input[:half])
	right := MergeSort(input[half:])
	return merge(left, right)
}

// merge function that merges two sorted slices into a single sorted slice
func merge(input1, input2 []int) []int {
	var newArray []int
	i, j := 0, 0

	for i < len(input1) && j < len(input2) {
		if input1[i] < input2[j] {
			newArray = append(newArray, input1[i])
			i++
		} else {
			newArray = append(newArray, input2[j])
			j++
		}
	}

	// Append remaining elements (if any) from input1
	for ; i < len(input1); i++ {
		newArray = append(newArray, input1[i])
	}

	// Append remaining elements (if any) from input2
	for ; j < len(input2); j++ {
		newArray = append(newArray, input2[j])
	}

	return newArray
}

func QuickSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	left, right := 0, len(input)-1

	// Pick a pivot
	pivotIndex := len(input) / 2

	// Move the pivot to the right
	input[pivotIndex], input[right] = input[right], input[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range input {
		if input[i] < input[right] {
			input[i], input[left] = input[left], input[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	input[left], input[right] = input[right], input[left]

	// Go down the rabbit hole
	QuickSort(input[:left])
	QuickSort(input[left+1:])

	return input
}
