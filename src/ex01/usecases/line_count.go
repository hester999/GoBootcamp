package usecases

import (
	"bufio"
	"os"
)

func LineCount(path string) (int, error) {

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, nil
}
