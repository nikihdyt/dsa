package solution

func RemDuplicates(nums []int) int {
	// make a map
	// check existensi
	mMap := make(map[int]bool)

	for _, num := range nums {
		_, exists := mMap[num]
		if exists { // if exists, next
			continue
		} else { // if not, add ke map
			mMap[num] = true
		}
	}

	return len(mMap)
}