## go-proverbs

> Go Proverbs for your code or command line.

[![Build Status](https://travis-ci.org/wayneashleyberry/go-proverbs.svg?branch=master)](https://travis-ci.org/wayneashleyberry/go-proverbs)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/go-proverbs)](https://goreportcard.com/report/github.com/wayneashleyberry/go-proverbs)
[![GoDoc](https://godoc.org/github.com/wayneashleyberry/go-proverbs?status.svg)](https://godoc.org/github.com/wayneashleyberry/go-proverbs)

### What are Go proverbs?

Check out [Rob Pike's video](https://www.youtube.com/watch?v=PAAkCSZUG1c).

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

### Motivation

I wanted a nice greeting message for my shell :)

![screen shot 2016-10-04 at 5 28 46 pm](https://cloud.githubusercontent.com/assets/727262/19080734/4708ab08-8a58-11e6-924b-60a350d0e926.png)
