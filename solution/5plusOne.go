package solution

func PlusOne(digits []int) []int { // failed in test case [9,9]
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
	}

	if digits[0] == 9 {
		digits = append([]int{0}, digits...)
	}

	for i, val := range digits {

		if val == 9 {
			digits[i-1] = digits[i-1] + 1
			digits[i] = 0
		}
	}

	return digits
}

func PlusOne2ndTry(digits []int) []int { // failed in test case [8,9,9,9] & [9,8,7,6,5,4,3,2,1,0]
	if digits[0] == 9 {
		digits = append([]int{0}, digits...)
	}

	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
	}

	for i, _ := range digits {
		if digits[len(digits)-i-1] == 9 || digits[len(digits)-i-1] == 10 {
			digits[len(digits)-i-1] = 0
			digits[len(digits)-i-2]++
		}
	}

	return digits
}

func PlusOne3rdTry(digits []int) []int { // failed in test case [9,8,7,6,5,4,3,2,1,0]
	if digits[0] == 9 {
		digits = append([]int{0}, digits...)
	}

	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}

	for i := 0; i < len(digits); i++ {
		if i == len(digits)-1 {
			continue
		}
		if digits[len(digits)-i-1] == 9 || digits[len(digits)-i-1] == 10 {
			digits[len(digits)-i-1] = 0
			digits[len(digits)-i-2]++
		}
	}

	return digits
}

func PlusOne4thTry(digits []int) []int { // failed in test case [9,8,9]
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}

	if digits[0] == 9 && digits[1] == 9 {
		digits = append([]int{0}, digits...)
	}

	for i := 0; i < len(digits); i++ {
		if i == len(digits)-1 {
			continue
		}
		if digits[len(digits)-i-1] == 9 || digits[len(digits)-i-1] == 10 {
			digits[len(digits)-i-1] = 0
			digits[len(digits)-i-2]++
		}
	}

	return digits
}

func PlusOne5thTry(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}
	
	for i := 0; i < len(digits); i++ { // digits[last] == 9  // i = index
		observedIdx := len(digits) - i - 1 // looping dari belakang ke depan

		digits[observedIdx]++ // digits[last_idx] pasti akan 10

		if digits[observedIdx] == 10 && observedIdx != 0 {
			digits[observedIdx] = 0
		} else {
			break
		}
	}

	if digits[0] == 10 {
		digits = append([]int{1}, digits...)
		digits[1] = 0
	}

	return digits
}