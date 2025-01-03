package solution

import "fmt"

func twoSum(nums []int, target int) [2]int {  // return: [0, 1]
	mMap := make(map[int]int)

	for idx, i := range nums { // i = 2
		complement := target - i // complement --> 7 - 2 = 5
		val, ok := mMap[complement] // ok --> ambil index (val)
		if ok {
			fmt.Println("oke gas oke gas")
			return [2]int{idx, val}
		} else {
			mMap[i] = idx
		}
		
		fmt.Println("i did this", i)
	}

	return [2]int{100, 100}
}

// func main() {
// 	a := []int{2, 5, 8, 4, 6}
// 	fmt.Println(twoSum(a, 7))
// }