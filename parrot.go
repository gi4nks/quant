package quant

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/ttacon/chalk"
)

func NewParrot(n string) *Parrot {
	return &Parrot{name: n, debugMode: false, on: true}
}

func NewVerboseParrot(n string) *Parrot {
	return &Parrot{name: n, debugMode: true, on: true}
}

type Parrot struct {
	name      string
	debugMode bool
	on        bool
}

func (t Parrot) trace(a Action0) {
	if t.on {
		a()
	}
}

func (t Parrot) Warn(a ...interface{}) {
	t.trace(func() {
		end := time.Now()

		fmt.Printf("{%s%v%s} |%s%s%s| ",
			chalk.Yellow, end.Format("2006/01/02 - 15:04:05"), chalk.Reset,
			chalk.Cyan, "WA", chalk.Reset)
		fmt.Println(a...)
	})
}

func (t Parrot) Error(a ...interface{}) {
	t.trace(func() {
		end := time.Now()

		fmt.Printf("{%s%v%s} |%s%s%s| ",
			chalk.Yellow, end.Format("2006/01/02 - 15:04:05"), chalk.Reset,
			chalk.Red, "ER", chalk.Reset)
		fmt.Println(a...)

	})
}

func (t Parrot) Info(a ...interface{}) {
	t.trace(func() {
		end := time.Now()

		fmt.Printf("{%s%v%s} |%s%s%s| ",
			chalk.Yellow, end.Format("2006/01/02 - 15:04:05"), chalk.Reset,
			chalk.White, "IN", chalk.Reset)
		fmt.Println(a...)
	})
}

func (t Parrot) Debug(a ...interface{}) {
	t.trace(
		func() {
			if t.debugMode {
				end := time.Now()

				fmt.Printf("{%s%v%s} |%s%s%s| ",
					chalk.Yellow, end.Format("2006/01/02 - 15:04:05"), chalk.Reset,
					chalk.Green, "DB", chalk.Reset)
				fmt.Println(a...)
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

func (t Parrot) TablePrint(header []string, body [][]string) {
	t.trace(
		func() {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader(header)
			table.SetBorder(false) // Set Border to false
			table.AppendBulk(body) // Add Bulk Data
			table.SetAlignment(3)
			table.Render()
		})
}
