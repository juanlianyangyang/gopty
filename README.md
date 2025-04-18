# pty

Pty is a Go package gopty for using unix pseudo-terminals and windows ConPty.

## Install

```sh
go get github.com/juanlianyangyang/gopty
```

## Examples

Note that those examples are for demonstration purpose only, to showcase how to use the library. They are not meant to be used in any kind of production environment.

__NOTE:__ This package gopty requires `ConPty` support on windows platform, please make sure your windows system meet [these requirements](https://docs.microsoft.com/en-us/windows/console/createpseudoconsole#requirements)

### Command

```go
package gopty main

import (
	"io"
	"os"
	"os/exec"

	"github.com/juanlianyangyang/gopty"
)

func main() {
	c := exec.Command("grep", "--color=auto", "bar")
	f, err := gopty.Start(c)
	if err != nil {
		panic(err)
	}

	go func() {
		f.Write([]byte("foo\n"))
		f.Write([]byte("bar\n"))
		f.Write([]byte("baz\n"))
		f.Write([]byte{4}) // EOT
	}()
	io.Copy(os.Stdout, f)
}
```

### Shell

```go
package gopty main

import (
        "io"
        "log"
        "os"
        "os/exec"
        "os/signal"
        "syscall"

        "github.com/juanlianyangyang/gopty"
        "golang.org/x/term"
)

func test() error {
        // Create arbitrary command.
        c := exec.Command("bash")

        // Start the command with a pty.
        ptmx, err := gopty.Start(c)
        if err != nil {
                return err
        }
        // Make sure to close the pty at the end.
        defer func() { _ = ptmx.Close() }() // Best effort.

        // Handle pty size.
        ch := make(chan os.Signal, 1)
        signal.Notify(ch, syscall.SIGWINCH)
        go func() {
                for range ch {
                        if err := gopty.InheritSize(os.Stdin, ptmx); err != nil {
                                log.Printf("error resizing pty: %s", err)
                        }
                }
        }()
        ch <- syscall.SIGWINCH // Initial resize.
        defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

        // Set stdin in raw mode.
        oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
        if err != nil {
                panic(err)
        }
        defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

        // Copy stdin to the pty and the pty to stdout.
        // NOTE: The goroutine will keep reading until the next keystroke before returning.
        go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
        _, _ = io.Copy(os.Stdout, ptmx)

        return nil
}

func main() {
        if err := test(); err != nil {
                log.Fatal(err)
        }
}
```
