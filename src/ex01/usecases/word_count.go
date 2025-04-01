package usecases

import (
	"bufio"
	"os"
	"unicode"
)

func WordCount(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	count := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		count += processWordLine(line)
	}

	return count, nil
}

func processWordLine(line string) int {
	isWord := false
	count := 0

	for _, r := range line {

		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			isWord = false
		} else if !isWord {
			isWord = true
			count++
		}
	}
	return count
}
