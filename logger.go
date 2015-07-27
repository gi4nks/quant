package quant

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

type logger interface {
	Print(v ...interface{})
}

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{log.New(os.Stdout, "\r\n", 0)}
}

// Format log
var sqlRegexp = regexp.MustCompile(`(\$\d+)|\?`)

func (logger Logger) Print(values ...interface{}) {
	currentTime := "\n\033[33m[" + time.Now().Format("2006-01-02 15:04:05") + "]\033[0m"
	source := fmt.Sprintf("\033[35m(%v)\033[0m", values[0])
	messages := []interface{}{source, currentTime}

	messages = append(messages, "\033[31;1m")
	messages = append(messages, values[0:]...)
	messages = append(messages, "\033[0m")
	
	logger.Println(messages...)
	
}