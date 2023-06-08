package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func StreamHandler(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var (
		name string
		sum  = 0.
		avg  = 0.
	)
	for i := 1; ; i++ {
		if !scanner.Scan() {
			break
		}
		if i == 1 {
			name = scanner.Text()
			continue
		}

		num, _ := strconv.Atoi(scanner.Text())

		sum += float64(num)
		avg = sum / float64(i)

		fmt.Printf("sfn received, name=%s, i=%d, num=%d, sum=%0.2f, avg=%0.2f\n", name, i, num, sum, avg)
	}
}
