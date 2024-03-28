package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yomorun/yomo"
)

func main() {
	sfn := yomo.NewStreamFunction("sfn", "localhost:9000")
	sfn.SetObserveDataTags(0x33, 0x34)
	err := sfn.Connect()
	if err != nil {
		log.Fatalf("[sfn] ❌ Connect to YoMo-Zipper failure with err: %v", err)
	}

	ch := make(chan os.Signal, 1)

	go func() {
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		fmt.Println("close sfn", time.Now())
		sfn.Close()
	}()

	sfn.Wait()
}
