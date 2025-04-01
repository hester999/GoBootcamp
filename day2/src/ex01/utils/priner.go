package utils

import (
	"fmt"
	"src/src/ex01/entity"
)

func PrintRes(res []entity.WorkerResult) {

	for _, v := range res {
		fmt.Println(v.Count, v.Name)
	}
}
