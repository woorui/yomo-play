package source

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

func PostToHttpSrv(name string, r io.Reader) error {
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
		_, err = http.Post("http://127.0.0.1:8090", "application/json", bytes.NewBuffer(b))
		if err != nil {
			return err
		}

	}
	return nil
}
