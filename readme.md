## go-proverbs

[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/go-proverbs)](https://goreportcard.com/report/github.com/wayneashleyberry/go-proverbs)
[![GoDoc](https://godoc.org/github.com/wayneashleyberry/go-proverbs?status.svg)](https://godoc.org/github.com/wayneashleyberry/go-proverbs)

### Package

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

### Command

```sh
go get github.com/wayneashleyberry/go-proverbs/...
```

```sh
goproverb
> Design the architecture, name the components, document the details.
```
