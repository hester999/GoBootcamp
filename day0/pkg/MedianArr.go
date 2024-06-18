package pkg

import "sort"

func MedianArr(arr []int) float64 {

	sort.Ints(arr)
	var res float64

	if len(arr)%2 != 0 {
		res = Mean(arr)
	} else {
		left := arr[0 : len(arr)/2]
		right := arr[(len(arr)/2)+1:]

		res = (Mean(left) + Mean(right)) / 2
	}

	return res
}
