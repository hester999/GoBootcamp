package main

import (
	_ "day1/ex00/pkg"
	pkg2 "day1/ex00/pkg"
	"fmt"
)

func main() {
	var f pkg2.Flag
	fmt.Println("hello world")
	s := f.ParseCommandLine()

	fmt.Println(s)

}

