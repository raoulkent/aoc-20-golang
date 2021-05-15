package days

import (
	"fmt"
	"sort"

	inputs "../inputs"
)

func One() {
	inputSlice := inputs.Day1
	var target = int(2020)
	sort.Ints(inputSlice[:])
	var a, b = WalkInward(inputSlice, target)
	var multiplied int = a*b
	fmt.Println(multiplied)
}

func WalkInward(arr []int, t int) (a, b int) {
	var lhs int = 0
	var rhs int = len(arr) - 1

	for rhs > lhs {
		if arr[rhs] > t {
			rhs--
		} else if arr[lhs] + arr[rhs] < t {
			lhs++
		} else if arr[lhs] + arr[rhs] > t {
			rhs--
		} else if arr[lhs] + arr[rhs] == t {
			return arr[lhs], arr[rhs]
		} else {
			panic("Oh no!")
		}
	}
	return
}
