package pkg

import (
	"sort"
)

func ModeArr(arr []int) int {

	sort.Ints(arr)

	m := make(map[int]int)

	for _, k := range arr {
		m[k]++
	}

	mode := 0
	frequency := 0

	for k, v := range m {
		if frequency < v {
			frequency = v
			mode = k
		} else if frequency == v {
			if k < mode {
				mode = k
			}
		}
	}
	return mode
}
