package quant

import (
	"fmt"
	"os"
)

func NewParrot(n string) *Parrot {
	return &Parrot{name: n, debugMode: false}
}

func NewVerboseParrot(n string) *Parrot {
	return &Parrot{name: n, debugMode: true}
}

type Parrot struct {
	name string
	debugMode bool
}

func (t Parrot) trace(a action0) {
	a()
}

func (t Parrot) Error(message string, err error) {
	t.trace(func() { fmt.Fprintln(os.Stderr, message, err) })
}

func (t Parrot) Info(message string) {
	t.trace(func() { fmt.Fprintln(os.Stdout, message) })
}

func (t Parrot) Debug(message string) {
	t.trace(
		func() { 
			if t.debugMode {
				fmt.Fprintln(os.Stdout, message) 
			}
		})
}

