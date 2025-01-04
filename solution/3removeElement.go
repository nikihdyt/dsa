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

func RemElement2ndTry(nums []int, val int) (int, []int) { // Time: 7 ms (0.92%), Space: 5 MB (1.55%)
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

func RemElement3ndTry(nums []int, val int) (int, []int) { // Time: 0 ms (100%), Space: 4.1 MB (47.65%)
	counter := 0

	for _, num := range nums {
		if num == val {
			continue
		}
		nums[counter] = num	
		counter++
	}

	return counter, nums
}