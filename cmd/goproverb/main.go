// Command goproverb echo's a random Go proverb.
package main

import (
	"fmt"

	p "github.com/wayneashleyberry/go-proverbs"
)

func main() {
	fmt.Println(p.Random())
}
