package main

import (
	"fmt"
	"sort"
)

type ByIntValue []int

func (a ByIntValue) Len() int {
	return a.Len()
}

func (a ByIntValue) Less(b int, c int) bool {
	return a[b] < a[c]
}

func (a ByIntValue) Swap(i int, j int) {
	i, j = j, i
}

type ByStringValue []string

func (a ByStringValue) Len() int {
	return a.Len()
}

func (a ByStringValue) Less(b int, c int) bool {
	return a[b] < a[c]
}

func (a ByStringValue) Swap(i int, j int) {
	i, j = j, i
}

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	// fmt.Println(xi)
	// // sort xi
	// fmt.Println(xi)

	// fmt.Println(xs)
	// // sort xs
	// fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	// sort.Sort(ByIntValue(xi))
	// sort.Sort(ByStringValue(xs))

	fmt.Println(xi)
	fmt.Println(xs)
}
