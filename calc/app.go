package main

import (
	"fmt"
	"sync"

	"github.com/tidwall/gjson"
	"github.com/yomorun/yomo/serverless"
)

type family struct {
	sum float64
	i   float64
	avg float64
}

var locker sync.Mutex

var familyMap = map[string]*family{}

// Handler will handle the raw data
func Handler(ctx serverless.Context) {
	data := ctx.Data()

	v := gjson.GetBytes(data, "Key")
	name := v.String()

	fmt.Println(name)

	v = gjson.GetBytes(data, "Value")
	num := v.Float()

	locker.Lock()
	defer locker.Unlock()

	f, ok := familyMap[name]
	if !ok {
		f = &family{}
		familyMap[name] = f
	}

	// calculate the average.
	f.sum += num
	f.avg = f.sum / f.i

	fmt.Printf("sfn received, name=%s, i=%0f, num=%0f, sum=%0.2f, avg=%0.2f\n", name, f.i, num, f.sum, f.avg)

	f.i++
}

func DataTags() []uint32 {
	return []uint32{12345}
}
