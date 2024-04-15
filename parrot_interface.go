package quant

type Parroter interface {
	Warn(a ...interface{})
	Error(a ...interface{})
	Info(a ...interface{})
	Debug(a ...interface{})
	Print(a ...interface{})
	Println(a ...interface{})
	Tablify(header []string, body [][]string)
}