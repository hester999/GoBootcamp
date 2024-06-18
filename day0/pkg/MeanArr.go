package pkg

func Mean(arr []int) float64 {

	sum := ArrSum(arr)
	res := sum / len(arr)

	return float64(res)
}
