package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 0)
	for index := 0; index < 10; index++ {
		a = append(a, index)
	}
	fmt.Println(a)

	index := sort.Search(len(a), func(_index int) bool {
		return a[_index] == 5
	})

	fmt.Println(a[index])
}
func searchInsert(nums []int, target int) int {
	sort.Search()
}
