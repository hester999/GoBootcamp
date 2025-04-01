package usecases

import (
	"fmt"
	"src/src/ex01/entity"
	"src/src/ex01/utils"
)

type CounterEx01 struct{}

func (c *CounterEx01) Process(flags interface{}) error {
	var process ProcessFunc
	var res []entity.WorkerResult
	flagSet, ok := flags.(*entity.Flags)
	if !ok {
		return fmt.Errorf("invalid flags type for CounterEx01")
	}
	if flagSet.LineCount {
		process = LineCount
		res = StartProcess(flagSet.Path, process)
	}
	if flagSet.WordCount {
		process = WordCount
		res = StartProcess(flagSet.Path, process)
	}
	if flagSet.CharCount {
		process = CharCount
		res = StartProcess(flagSet.Path, process)
	}

	utils.PrintRes(res)

	return nil
}
