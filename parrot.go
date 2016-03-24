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

func (t Parrot) Warn(a ...interface{}) {
	t.trace(func() {
		fmt.Print(chalk.Yellow)
		fmt.Println(a...)
		fmt.Print(chalk.Reset)
	})
}

func (t Parrot) Error(a ...interface{}) {
	t.trace(func() {
		fmt.Print(chalk.Red)
		fmt.Println(a...)
		fmt.Print(chalk.Reset)

	})
}

func (t Parrot) Info(a ...interface{}) {
	t.trace(func() {
		fmt.Print(chalk.White)
		fmt.Println(a...)
		fmt.Print(chalk.Reset)
	})
}

func (t Parrot) Debug(a ...interface{}) {
	t.trace(
		func() {
			if t.debugMode {
				fmt.Print(chalk.Green)
				fmt.Println(a...)
				fmt.Print(chalk.Reset)
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
