package source

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/yomorun/yomo"
)

const DataTag = 12345

type KV struct {
	Key   string
	Value string
}

// PipeToSource is the main logic of source.
func PipeToSource(name string, r io.Reader, source yomo.Source) error {
	scanner := bufio.NewScanner(r)

	for i := 1; ; i++ {
		if !scanner.Scan() {
			break
		}
		fmt.Println(name, scanner.Text())
		b, err := json.Marshal(&KV{Key: name, Value: scanner.Text()})
		if err != nil {
			return err
		}
		if err := source.Write(DataTag, b); err != nil {
			return err
		}
	}
	return nil
}
