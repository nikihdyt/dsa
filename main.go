package main

import (
	"fmt"
	"main/solution"
)

func main() {
	a := []int{4,3,2,1}
	b := []int{1,3,2,9}
	c := []int{9,9}
	d := []int{8,9,9,9}
	e := []int{9,8,7,6,5,4,3,2,1,0}
	f := []int{9,8,9}
	g := []int{9,8,9}
	fmt.Println(solution.PlusOne5thTry(a))
	fmt.Println(solution.PlusOne5thTry(b))
	fmt.Println(solution.PlusOne5thTry(c))
	fmt.Println(solution.PlusOne5thTry(d))
	fmt.Println(solution.PlusOne5thTry(e))
	println("--- test case [9,8,9] ---")
	fmt.Println(solution.PlusOne4thTry(f))
	fmt.Println(solution.PlusOne5thTry(g))
}