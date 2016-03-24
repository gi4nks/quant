package quant

import (
	"fmt"
	//"os"
	"time"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
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

func colorForStatus(method string) string {
	switch {
	case method == "debug":
		return green
	case method == "info":
		return white
	case method == "warn":
		return yellow
	default:
		return red
	}
}

func (t Parrot) trace(a Action0) {
	a()
}

func (t Parrot) Warn(a ...interface{}) {
	t.trace(func() {
		end := time.Now()
		statusColor := colorForStatus("warn")

		fmt.Printf("[%v] |%s %s %s| ",
			end.Format("2006/01/02 - 15:04:05"),
			statusColor, "WA", reset)
		fmt.Println(a...)
	})
}

func (t Parrot) Error(a ...interface{}) {
	t.trace(func() {
		end := time.Now()
		statusColor := colorForStatus("error")

		fmt.Printf("[%v] |%s %s %s| ",
			end.Format("2006/01/02 - 15:04:05"),
			statusColor, "ER", reset)
		fmt.Println(a...)

	})
}

func (t Parrot) Info(a ...interface{}) {
	t.trace(func() {
		end := time.Now()
		statusColor := colorForStatus("info")

		fmt.Printf("[%v] |%s %s %s| ",
			end.Format("2006/01/02 - 15:04:05"),
			statusColor, "IN", reset)
		fmt.Println(a...)
	})
}

func (t Parrot) Debug(a ...interface{}) {
	t.trace(
		func() {
			if t.debugMode {
				end := time.Now()
				statusColor := colorForStatus("debug")

				fmt.Printf("[%v] |%s %s %s| ",
					end.Format("2006/01/02 - 15:04:05"),
					statusColor, "DG", reset)
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
