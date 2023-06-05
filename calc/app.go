package main

import (
	"fmt"
	"strconv"

	"github.com/yomorun/yomo/serverless"
)

var (
	sum = 0.
	i   = float64(1)
)

// Handler will handle the raw data
func Handler(ctx serverless.Context) {
	data := ctx.Data()

	num, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Printf("sfn received NaN: %+v\n", data)
		return
	}

	// calculate the average.
	sum += float64(num)
	avg := sum / i

	fmt.Printf("sfn received, i=%0f, num=%d, sum=%0.2f, avg=%0.2f\n", i, num, sum, avg)

	i++
}

func DataTags() []uint32 {
	return []uint32{12345}
}
