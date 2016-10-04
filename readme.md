## go-proverbs

> Go Proverbs for your code or command line.

[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/go-proverbs)](https://goreportcard.com/report/github.com/wayneashleyberry/go-proverbs)
[![GoDoc](https://godoc.org/github.com/wayneashleyberry/go-proverbs?status.svg)](https://godoc.org/github.com/wayneashleyberry/go-proverbs)

### What are Go proverbs?

Check out [Rob Pike's video](https://www.youtube.com/watch?v=PAAkCSZUG1c).

### Motivation

### Usage

#### As a package

```sh
go get github.com/wayneashleyberry/go-proverbs/
```

```go
package main

import (
	"fmt"

	p "github.com/wayneashleyberry/go-proverbs"
)

func main() {
    // Prints a random Go proverb.
	fmt.Println(p.Random())
}
```

#### As a command

```sh
go get github.com/wayneashleyberry/go-proverbs/...
```

```sh
goproverb
> Design the architecture, name the components, document the details.
```
