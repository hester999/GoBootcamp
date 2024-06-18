package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func InputNum() []int {

	sc := bufio.NewScanner(os.Stdin)

	NumArr := []int{}
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if num < -100000 || num > 100000 {
			fmt.Fprintln(os.Stderr, "Input number is out of range")
		}
		NumArr = append(NumArr, num)
	}
	return NumArr
}
