package quant

import (
	"fmt"
	"os"
)

func NewParrot(n string) *Parrot {
	return &Parrot{name: n}
}

type Parrot struct {
	name string
}

func (t Parrot) trace(a action) {
	a()
}

func (t Parrot) Error(message string, err error) {
	t.trace(func() { fmt.Fprintln(os.Stderr, message, err) })
}

func (t Parrot) Info(message string) {
	t.trace(func() { fmt.Fprintln(os.Stdout, message) })
}
