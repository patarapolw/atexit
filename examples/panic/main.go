package main

import (
	"fmt"
	"time"

	"github.com/patarapolw/atexit"
)

func handler() {
	fmt.Println("atexit triggered")
}

func main() {
	atexit.Listen()
	defer atexit.ListenPanic()

	atexit.Register(handler)

	go func() {
		defer atexit.ListenPanic()

		time.Sleep(1 * time.Second)
		panic("panic")
	}()

	time.Sleep(1 * time.Minute)
	atexit.Exit(0)
}
