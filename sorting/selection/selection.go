package selection

func Sort(array []int) []int {
	for i := 0; i < len(array); i++ {
		actualValue := array[i]
		lessIdx := i

		for j := i + 1; j < len(array); j++ {
			nextValue := array[j]
			lessValue := array[lessIdx]

			if nextValue < lessValue {
				lessIdx = j
			}
		}

		if lessIdx > 0 {
			array[i] = array[lessIdx]
			array[lessIdx] = actualValue
		}
	}

	return array
}
