//go:build js

package jlog

import (
	"fmt"
	"io"

	"syscall/js"
)

var PackageLevel = LevelNone

type Level int

const (
	LevelNone Level = iota
	LevelDebug
	LevelTrace
)

func print(a ...interface{}) {
	fmt.Fprintln(Console, a...)
}

func printf(format string, a ...interface{}) {
	fmt.Fprintf(Console, format, a...)
}

// Debug print
func Debug(a ...interface{}) {
	if PackageLevel >= LevelDebug {
		print(a...)
	}
}

func Debugf(format string, a ...interface{}) {
	if PackageLevel >= LevelDebug {
		printf(format, a...)
	}
}

func Trace(a ...interface{}) {
	if PackageLevel >= LevelTrace {
		print(a...)
	}
}

func Tracef(format string, a ...interface{}) {
	if PackageLevel >= LevelTrace {
		printf(format, a...)
	}
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
