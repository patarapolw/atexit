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
	atexit.Register(handler)
	atexit.Listen()

	go func() {
		defer atexit.ListenPanic()

		time.Sleep(1 * time.Second)
		s := []int{}
		s[0] = s[0] + 1
	}()

	time.Sleep(1 * time.Minute)
	atexit.Exit(0)
}
