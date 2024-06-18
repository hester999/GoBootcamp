package main

import (
	"day0/pkg"
	"flag"
	"fmt"
)

type Flag struct {
	mean   bool
	median bool
	mod    bool
	sd     bool
}

func (f *Flag) ParseFlags() {
	// Определяем флаги
	flag.BoolVar(&f.mean, "mean", false, "Calculate and display the mean")
	flag.BoolVar(&f.median, "median", false, "Calculate and display the median")
	flag.BoolVar(&f.mod, "mod", false, "Calculate and display the mode")
	flag.BoolVar(&f.sd, "sd", false, "Calculate and display the standard deviation")

	// Парсим флаги
	flag.Parse()
}
func (f *Flag) Mean(arr []int) {
	res := pkg.Mean(arr)
	fmt.Printf("Mean: %.2f\n", res)
}
func (f *Flag) Median(arr []int) {
	res := pkg.MedianArr(arr)
	fmt.Printf("Median: %.2f\n", res)
}

func (f *Flag) Mod(arr []int) {
	res := pkg.ModeArr(arr)
	fmt.Printf("Mod: %d\n", res)
}

func (f *Flag) Sd(arr []int) {
	res := pkg.StandardDeviation(arr)
	fmt.Printf("Sd: %.2f\n", res)
}
func main() {
	f := Flag{}
	f.ParseFlags()
	arr := pkg.InputNum()

	// Проверяем и выводим, какие флаги были установлены
	if f.mean {
		f.Mean(arr)
	}
	if f.median {
		f.Median(arr)
	}
	if f.mod {
		f.Mod(arr)
	}
	if f.sd {
		f.Sd(arr)
	}
	if !f.mean && !f.median && !f.mod && !f.sd {
		// Если ни один флаг не установлен, выполняем все
		f.Mean(arr)
		f.Median(arr)
		f.Mod(arr)
		f.Sd(arr)
	}
}
