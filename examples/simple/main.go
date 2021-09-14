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
	atexit.Listen() // Await for SIGINT, SIGTERM, whatever. Also works in Windows

	time.Sleep(1 * time.Minute)
	atexit.Exit(0)
}
