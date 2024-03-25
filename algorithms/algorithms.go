package algorithms

import "sorting_algorithms/sorting/selection"

func SymetricDifference(arrays ...[]int) []int {
	result := symetricComparisson(arrays[0], arrays[1])

	if len(arrays) > 2 {
		for i := 2; i < len(arrays); i++ {
			result = symetricComparisson(arrays[i], result)
		}
	}

	result = selection.Sort(result)
	result = uniqueItems(result)

	return result
}

func symetricComparisson(arrays ...[]int) []int {
	result := []int{}

	for i, arr := range arrays {
		for j, nexArr := range arrays {
			if i == j {
				continue
			}

		mainLoop:
			for _, v1 := range arr {
				for _, v2 := range nexArr {
					if v1 == v2 {
						continue mainLoop
					}
				}

				result = append(result, v1)
			}
		}
	}

	return result
}

func uniqueItems(slice []int) []int {
	tmpMap := map[int]int{}
	cleanSlice := []int{}

	for _, v := range slice {
		_, exists := tmpMap[v]
		if exists {
			continue
		}

		cleanSlice = append(cleanSlice, v)
		tmpMap[v] = v
	}

	return cleanSlice
}

func CallupdateInventory(n, r int) int {
	nFact := factorial(n)
	res := factorial(n - r)

	return nFact / res
}

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * factorial(number-1)
}

func updateInventory(arr1, arr2 [][]interface{}) [][]interface{} {

	return nil
}

func Pairwise(arr []int, arg int) int {
	usedFlags := map[int]bool{}
	result := 0

	for i, v := range arr {
		_, exist := usedFlags[i]
		if exist {
			// this number has been used no need to compare with others
			continue
		}

		for j, v2 := range arr {
			if i == j {
				// do not compare to itself
				continue
			}
			_, exist := usedFlags[j]
			if exist {
				continue
			}

			if v+v2 == arg {
				usedFlags[i] = true
				usedFlags[j] = true
				result += i + j
				break
			}
		}
	}

	return result
}
