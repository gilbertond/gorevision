package main

/*
Infinite array, all valid values from the beginning, junk towards the end, return index of valid number and if non found return -1
*/

func main() {
	arr := make([]interface{}, 0)
	// arr = append(arr, 1, "a")
	arr = append(arr, 1, 2, 3, 4, 10, 20, nil, "ok")
	// arr = append(arr, 1, 2, 3, nil, 4, "ok")

	println(findIndex(arr, 3))
}

func findIndex(infiniteArray []interface{}, value int) int {
	if len(infiniteArray) == 0 {
		return -1
	}

	sIndex := 2
	if isNotValid(infiniteArray[sIndex]) {
		sIndex = sIndex / 2
	}

	if !isNotValid(infiniteArray[sIndex]) {
		sIndex = sIndex * 2
	}

	return -1
}

func findIndexV2(infiniteArray []interface{}, value int) int {
	if len(infiniteArray) == 0 {
		return -1
	}
	for i := 0; i < len(infiniteArray); {

		if !isNotValid(infiniteArray[i]) {
			return -1
		}

		if infiniteArray[i] == value {
			return i
		}

		i++
	}
	return -1
}

func isNotValid(someValue interface{}) bool {
	if _, ok := someValue.(int); ok {
		return true
	}
	return false
}
