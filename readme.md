![Screenshot](https://cloud.githubusercontent.com/assets/727262/26555544/c7b40ce6-4495-11e7-8d48-28100d071b22.png)

> Go Proverbs as a package and cli.

[![Build Status](https://travis-ci.org/wayneashleyberry/go-proverbs.svg?branch=master)](https://travis-ci.org/wayneashleyberry/go-proverbs)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/go-proverbs)](https://goreportcard.com/report/github.com/wayneashleyberry/go-proverbs)
[![GoDoc](https://godoc.org/github.com/wayneashleyberry/go-proverbs?status.svg)](https://godoc.org/github.com/wayneashleyberry/go-proverbs)

### What are Go proverbs?

Check out [Rob Pike's](https://twitter.com/rob_pike) inspiring [talk at Gopherfest SV 2015](https://www.youtube.com/watch?v=PAAkCSZUG1c).

### Usage

#### CLI

```sh
go get github.com/wayneashleyberry/go-proverbs/...
```

```sh
$ goproverb
Design the architecture, name the components, document the details.
```


#### Package

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

### License

MIT © [Wayne Ashley Berry](https://wayne.cloud)
