package source

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yomorun/yomo"
)

const DataTag = 12345

// PipeToSource is the main logic of source.
func PipeToSource(r io.Reader, source yomo.Source) error {
	scanner := bufio.NewScanner(r)

	for i := 1; ; i++ {
		if !scanner.Scan() {
			break
		}
		fmt.Println(scanner.Text())
		if err := source.Write(DataTag, scanner.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
