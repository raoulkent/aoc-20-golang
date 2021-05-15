package days

import (
	"fmt"
	"sort"

	inputs "../inputs"
)

func One() {
	inputSlice := inputs.Day1
	sort.Ints(inputSlice[:])
	var target = int(2020)

	var a, b = walkInward(inputSlice, target)
	fmt.Println(a*b)
}

func walkInward(arr []int, t int) (a, b int) {
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
