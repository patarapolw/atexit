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
	atexit.Listen() // Await for SIGINT, SIGTERM, whatever. Also works in Windows
	defer atexit.ListenPanic() // Listen for panic and crashes

	atexit.Register(handler)

	go func() {
		defer atexit.ListenPanic()

		time.Sleep(1 * time.Second)
		panic("panic")
	}()
	time.Sleep(1 * time.Minute)

	atexit.Exit(0) // This also needs to be called at the end of main function, if you want atexit to be executed on normal exit.
}
```

## Caveats

- `os.Exit` doesn't call atexit. Use `atexit.Exit` instead.
- `atexit.Exit(0)` needs to be called at the end of main function, if you want atexit to be executed on normal exit.
- `log.Fatal*` also don't call atexit. Use `atexit.Fatal*` instead.
- `panic` can be listened, but do call `defer atexit.ListenPanic()` at the beginning of main function.
- If `panic` is inside Goroutine, call `defer atexit.ListenPanic()` at the beginning of main function, too.

## Install

    go get github.com/patarapolw/atexit
