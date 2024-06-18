package pkg

import "math"

func StandardDeviation(arr []int) float64 {

	mean := Mean(arr)

	sum := 0.0

	for _, k := range arr {

		sum += math.Pow(float64(k)-mean, 2)
	}

	res := sum / float64(len(arr))
	res = math.Sqrt(res)
	return res

}
