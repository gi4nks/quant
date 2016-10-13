package quant

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

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
			// setup the tab writer
			w := tabwriter.NewWriter(os.Stdout, 20, 1, 3, ' ', 0)

			h := strings.Join(header, "\t")

			// print header
			fmt.Fprintln(w, h)

			for _, p := range body {
				b := strings.Join(p, "\t")
				bb := b + "\n"
				fmt.Fprintf(w, bb)
			}

			w.Flush()

		})
}
