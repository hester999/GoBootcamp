package main

import (
	"fmt"
	"src/src/ex00/entity"
	"src/src/ex00/usecases"
	"src/src/external"
)

func main() {
	flags := &entity.Flags{}
	flags.Parse()

	var finder external.Finder = &usecases.FinderEx00{}
	if err := finder.Process(flags); err != nil {
		fmt.Println(err)
	}
}
