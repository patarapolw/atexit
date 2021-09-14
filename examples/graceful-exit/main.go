package main

import (
	"fmt"

	"github.com/patarapolw/atexit"
)

func handler() {
	fmt.Println("Exiting")
}

func main() {
	atexit.Register(handler)
	atexit.Exit(0)
}
