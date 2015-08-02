package quant

import (
	"errors"
	"github.com/op/go-logging"
	"os"
)

const (
	format_full = "%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}"
	format_light = "%{color}%{message}"
)

func NewTrace(n string) *Trace {
	return &Trace{name: n}
}

type Trace struct {
	name string
	logger *logging.Logger
}

func (t *Trace) init(format_name string) {
	t.logger= logging.MustGetLogger(t.name)
	format := logging.MustStringFormatter(format_name)
	
	// Configuring logger
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	// Set the backends to be used.
	logging.SetBackend(backendFormatter)
	// -------------------
}

func (t *Trace) Full() {
	t.init(format_full)
}

func (t *Trace) Light() {
	t.init(format_light)
}

func (t Trace) trace(a action0) error {
	if (t.logger == nil) {
		return errors.New("tracer has not bee corretcly initialized, please very to call Full or Emply function before use")
	}
	
	a()
	return nil
}


func (t Trace) Error(message string) error {
	return t.trace(func() { t.logger.Error(message) })
}

func (t Trace) Warning(message string) error {
	return t.trace(func() { t.logger.Warning(message) })
}

func (t Trace) Notice(message string) error {
	return t.trace(func() { t.logger.Info(message) })
}

func (t Trace) News(message string) error {
	return t.trace(func() { t.logger.Debug(message) })
}