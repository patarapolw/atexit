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
	atexit.Listen()            // Await for SIGINT, SIGTERM, whatever. Also works in Windows
	defer atexit.ListenPanic() // Listen for panic and crashes

	atexit.Register(handler)
	time.Sleep(1 * time.Minute)

	atexit.Exit(0) // This also needs to be called at the end of main function, if you want atexit to be executed on normal exit.
}
