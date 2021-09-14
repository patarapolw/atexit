# atexit

Simple `atexit` implementation for [Go](https://golang.org).

Note that you *have* to call `atexit.Exit` and not `os.Exit` to terminate your
program (that is, if you want the `atexit` handlers to execute).

## Example usage

```go
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
	atexit.Exit(0) // This also needs to be called at the end of main function.
}
```

## Caveats

- `os.Exit` won't work. Use `atexit.Exit` instead.
- `atexit.Exit(0)` needs to be called at the end of main function.
- `log.Fatal*` also don't work. Use `atexit.Fatal*` instead.
- `panic` without `recover` won't be listened. See [/examples/panic/main.go](/examples/panic/main.go)

## Install

    go get github.com/patarapolw/atexit
