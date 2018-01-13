package debug

import "fmt"

type debugger struct {
	debug bool
}

func NewDebugger(debug bool) *debugger {
	return &debugger{debug: debug}
}

func (d *debugger) Log(message ...interface{}) {
	if d.debug {
		fmt.Println(message)
	}
}

func (d *debugger) Logf(format string, args ...interface{}) {
	if d.debug {
		fmt.Printf(format, args)
	}
}
