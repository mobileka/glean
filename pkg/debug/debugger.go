package debug

type Debugger interface {
	Log(message string)
	Logf(format string, args ...interface{})
}
