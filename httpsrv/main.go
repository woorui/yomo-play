package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/tidwall/gjson"
)

type family struct {
	sum float64
	i   float64
	avg float64
}

var locker sync.Mutex

var familyMap = map[string]*family{}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// read data is diff way.
		data, _ := io.ReadAll(r.Body)

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
	})

	http.ListenAndServe(":8090", mux)
}
