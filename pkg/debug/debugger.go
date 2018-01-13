package debug

type Debugger interface {
	Log(message ...interface{})
	Logf(format string, args ...interface{})
}
