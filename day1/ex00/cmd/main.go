package main

import (
	pkg2 "day1/ex00/pkg"
	"fmt"
)

func main() {
	flag := pkg2.Flag{}
	str := flag.ParseCommandLine()
	recipes, err := pkg2.GetDB(str)
	if err != nil {
		fmt.Println("Error:", err)
	}

	pkg2.PrintDB(recipes, str)
}
