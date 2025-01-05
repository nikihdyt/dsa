package solution

// input: sorted nums & target
// return which idx is the target placed
// if there's no target in the nums, return the index of where it should be placed
// must write in O(log n)
func SearchIdx(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high{
		middle := low + ((high-low)/2)

		if nums[middle] == target {
			return middle
		}

		if target < nums[middle] {
			high = middle // high = middle - 1 // bingung
		} else {
			low = middle // low = middle + 1 // bingung
		}
	}
	
	return low // high // bingung
}