package main

import (
	"fmt"
	"src/src/ex01/entity"
	"src/src/ex01/usecases"
)

func main() {

	flags := entity.Flags{}

	err := flags.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	counter := usecases.CounterEx01{}

	err = counter.Process(&flags)
	if err != nil {
		fmt.Println(err)
		return
	}

}
