package quant

import (
	"fmt"
	//"os"
	"github.com/ttacon/chalk"
)

func NewParrot(n string) *Parrot {
	return &Parrot{name: n, debugMode: false}
}

func NewVerboseParrot(n string) *Parrot {
	return &Parrot{name: n, debugMode: true}
}

type Parrot struct {
	name      string
	debugMode bool
}

func (t Parrot) trace(a Action0) {
	a()
}

func (t Parrot) Warn(message string, err error) {
	//t.trace(func() { fmt.Fprintln(os.Stderr, message, err) })
	t.trace(func() { fmt.Println(chalk.Yellow, message, err, chalk.Reset) })
}

func (t Parrot) Error(message string, err error) {
	//t.trace(func() { fmt.Fprintln(os.Stderr, message, err) })
	t.trace(func() { fmt.Println(chalk.Red, message, err, chalk.Reset) })
}

func (t Parrot) Info(message string) {
	//t.trace(func() { fmt.Fprintln(os.Stdout, message) })
	t.trace(func() { fmt.Println(chalk.White, message, chalk.Reset) })
}

func (t Parrot) Debug(message string) {
	t.trace(
		func() {
			if t.debugMode {
				//fmt.Fprintln(os.Stdout, message)
				fmt.Println(chalk.Green, message, chalk.Reset)
			}
		})
}

func (t Parrot) Print(a ...interface{}) {
	t.trace(
		func() {
			fmt.Print(a...)
		})
}

func (t Parrot) Println(a ...interface{}) {
	t.trace(
		func() {
			fmt.Println(a...)
		})
}
