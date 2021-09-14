# About

Simple `atexit` implementation for [Go](https://golang.org).

Note that you *have* to call `atexit.Exit` and not `os.Exit` to terminate your
program (that is, if you want the `atexit` handlers to execute).

# Example usage

```go
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
```

# Install

    go get github.com/patarapolw/atexit
