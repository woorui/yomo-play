package main

import (
	"fmt"

	"github.com/yomorun/yomo/serverless"
)

// Init will initialize the stream function
func Init() error {
	fmt.Println("sfn init")
	return nil
}

// Handler will handle the raw data
func Handler(ctx serverless.Context) {
	data := ctx.Data()
	fmt.Printf("<< sfn received[%d tag, %d Bytes]: %s\n", ctx.Tag(), len(data), string(data))
	ctx.WriteWithTarget(0x35, data, "my-backflow")
}

func DataTags() []uint32 {
	return []uint32{0x33}
}

func WantedTarget() string {
	return "my-handler"
}
