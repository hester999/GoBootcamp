package usecases

import (
	"src/src/ex01/entity"
	"src/src/ex01/utils"
	"sync"
)

type ProcessFunc func(string) (int, error)

func StartProcess(files []string, processFunc ProcessFunc) []entity.WorkerResult {
	const maxWorker = 3
	inChan := make(chan string, 3)
	outChan := make(chan entity.WorkerResult, 3)
	wg := sync.WaitGroup{}

	go func() {
		defer close(inChan)
		for _, file := range files {
			inChan <- file
		}
	}()

	for i := 1; i <= maxWorker; i++ {
		wg.Add(1)
		go worker(inChan, outChan, &wg, processFunc)

	}

	go func() {
		wg.Wait()
		close(outChan)
	}()

	var results []entity.WorkerResult
	for out := range outChan {
		results = append(results, out)
	}

	return results
}

func worker(in <-chan string, out chan<- entity.WorkerResult, wg *sync.WaitGroup, processFunc ProcessFunc) {
	defer wg.Done()
	var err error
	for file := range in {
		var result entity.WorkerResult
		result.Count, err = processFunc(file)
		if err != nil {
			panic(err)
		}
		result.Name = utils.SearchName(file)
		out <- result
	}
}
