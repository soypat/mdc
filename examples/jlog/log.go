//go:build js

package jlog

import (
	"fmt"
	"io"

	"syscall/js"
)

func print(a ...interface{}) {
	if a == nil {
		return
	}
	fmt.Fprintln(Console, a...)
}

// Debug print
func Debug(a ...interface{}) {
	print(a...)
}

func Trace(a ...interface{}) {
	print(a...)
}

var Console io.Writer = jsWriter{
	Value: js.Global().Get("console"),
	fname: "log",
}

type jsWriter struct {
	js.Value
	fname string
}

func (j jsWriter) Write(b []byte) (int, error) {
	j.Call(j.fname, string(b))
	return len(b), nil
}
