package usecases

import (
	"bufio"
	"os"
	"unicode"
)

func CharCount(path string) (int, error) {
	file, err := os.Open(path)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count int

	for scanner.Scan() {
		line := scanner.Text()

		count += processCharLine(line)
	}
	return count, nil
}

func processCharLine(line string) int {

	count := 0

	for _, r := range line {
		if !unicode.IsControl(r) {
			count++
		}
	}
	return count
}
