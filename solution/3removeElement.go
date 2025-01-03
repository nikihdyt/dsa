package solution

import "fmt"

func RemElement(nums []int, val int) (int, []int) { // error: index out of range
	// val pointer

	// looping nums
	// ketemu val --> nums[idx] = nums [pointer + 1]
	// pointer ++
	// return pointer
	counter := 0

	for idx, num := range nums {
		if num == val {
			nums[idx] = nums[counter+1]
			fmt.Println("idx--- ", idx)
		} else {
			counter++
		}
	}

	return counter, nums
}

func RemElement2ndTry(nums []int, val int) (int, []int) {
	// yg di execute negative if nya

	// if num != val 
	// --> masukkan value itu ke index yg seharusnya
	counter := 0

	for _, num := range nums {
		if num != val {
			nums[counter] = num	
			counter++
		}
		fmt.Println(nums)
	}

	return counter, nums
}