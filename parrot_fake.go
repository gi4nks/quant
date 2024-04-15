package quant

import (
	"fmt"
)

// FakeParrot is a fake implementation of the Parrot struct for testing purposes.
type FakeParrot struct {
	name      string
	debugMode bool
	on        bool
}

// NewFakeParrot creates a new instance of FakeParrot.
func NewFakeParrot(n string) *FakeParrot {
	return &FakeParrot{name: n, debugMode: false, on: true}
}

// Warn prints warning messages.
func (p *FakeParrot) Warn(a ...interface{}) {
	fmt.Print("[WARN]: ")
	fmt.Println(a...)
}

// Error prints error messages.
func (p *FakeParrot) Error(a ...interface{}) {
	fmt.Print("[ERROR]: ")
	fmt.Println(a...)
}

// Info prints informational messages.
func (p *FakeParrot) Info(a ...interface{}) {
	fmt.Print("[INFO]: ")
	fmt.Println(a...)
}

// Debug prints debug messages.
func (p *FakeParrot) Debug(a ...interface{}) {
	fmt.Print("[DEBUG]: ")
	fmt.Println(a...)
}

// Println prints messages with a newline.
func (p *FakeParrot) Println(a ...interface{}) {
	fmt.Println(a...)
}

// Print prints messages without formatting.
func (p *FakeParrot) Print(a ...interface{}) {
	fmt.Print(a...)
}

// Tablify simulates tabular output.
func (p *FakeParrot) Tablify(header []string, body [][]string) {
	fmt.Println("Header:", header)
	fmt.Println("Body:", body)
}
