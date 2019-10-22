// Package goproverbs provides conveniant methods for Go proverbs. The proverbs
// were taken from https://go-proverbs.github.io/ and were originally inspired
// by Rob Pike's talk which can be seen here: https://youtu.be/PAAkCSZUG1c
package goproverbs

import (
	"math/rand"
	"time"
)

var proverbs = []string{
	"A little copying is better than a little dependency.",
	"Cgo is not Go.",
	"Cgo must always be guarded with build tags.",
	"Channels orchestrate; mutexes serialize.",
	"Clear is better than clever.",
	"Concurrency is not parallelism.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't communicate by sharing memory, share memory by communicating.",
	"Don't just check errors, handle them gracefully.",
	"Don't panic.",
	"Errors are values.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"interface{} says nothing.",
	"Make the zero value useful.",
	"Reflection is never clear.",
	"Syscall must always be guarded with build tags.",
	"The bigger the interface, the weaker the abstraction.",
	"With the unsafe package there are no guarantees.",
}

// All returns all Go proverbs
func All() []string {
	return proverbs
}

// Random returns a random Go proverb
func Random() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	i := r.Intn(len(proverbs))
	return proverbs[i]
}

// First returns the first Go proverb (alphabetically)
func First() string {
	return proverbs[0]
}

// Last returns the last Go proverb (alphabetically)
func Last() string {
	return proverbs[len(proverbs)-1]
}
