package goproverbs

import (
	"math/rand"
	"time"
)

var proverbs []string = []string{
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
	"Make the zero value useful.",
	"Reflection is never clear.",
	"Syscall must always be guarded with build tags.",
	"The bigger the interface, the weaker the abstraction.",
	"With the unsafe package there are no guarantees.",
	"interface{} says nothing.",
}

func All() []string {
	return proverbs
}

func Random() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	i := r.Intn(len(proverbs))
	return proverbs[i]
}

func First() string {
	return proverbs[0]
}

func Last() string {
	return proverbs[len(proverbs)]
}
