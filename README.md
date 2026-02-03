# got-style

The `got-style` package provides a simple and semantic method to style terminal outputs in go. Under the hood, the package utilizes ANSI style modes to format but encapsulates them into a user-friendly API.

## Basic Usage

See package docs for full documentation.

Use the style constants to create new style objects. Then chain with the built in print or format functions to display to terminal.

```go
import "github.com/binarysoupdev/got-style/style"

func main() {
    style.New(style.BOLD, style.GREEN).Println("Success Message!")
}
```

Several pre-built styles are also included for convenience.

```go
import "github.com/binarysoupdev/got-style/style"

func main() {
    style.BoldSuccess.Println("Success Message!")
}
```
