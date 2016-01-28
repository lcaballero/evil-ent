package ent

import "io"

type ExitStrategy func()

// A LogWriter is a super-interface for all the writing an application might
// attempt, so that instead of importing fmt, log, somelib.log, etc, a single
// interface can be used and injected into application code.  At which point
// imports can be scanned to guarantee that the application code doesn't use
// any unwanted log or printing imports.
type LogWriter interface {
	LevelWriter
	ConsoleWriter
	FatalWriter
	PanicWriter
	ErrorFactory
	Flusher
	Closer
}

type Flusher interface {
	Flush()
}

type Closer interface {
	io.Closer
	Closed() bool
}

// Provides methods to writing to a configured log sink limited to these
// log levels.
type LevelWriter interface {

	// Logs the values provided with level = Error
	Error(v ...interface{})

	// Logs the interpolated value with level = Error
	Errorf(format string, v ...interface{})

	// Logs the values provided with level = Debug
	Debug(v ...interface{})

	// Logs the interpolated value with level = Debug
	Debugf(format string, v ...interface{})

	// Logs the values provided with level = Warn
	Warn(v ...interface{})

	// Logs the interpolated value with level = Warn
	Warnf(format string, v ...interface{})

	// Logs the values provided with level = Info
	Info(v ...interface{})

	// Logs the interpolated value with level = Info
	Infof(format string, v ...interface{})

	// Logs the values provided with level = Trace
	Trace(v ...interface{})

	// Logs the interpolated values with level = Trace
	Tracef(format string, v ...interface{})
}

// Functions to print to Stdout.  These calls do not adhere to a log level.
// They are intended to force output to console for purposes of debugging
// regardless of the current level of logging, which could be 'off'.
type ConsoleWriter interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

// Part of handling truly dire situations.  These methods attempt to
// log and then bail from the running process immediately.
type FatalWriter interface {
	// Equivalent to Error() followed by an os.Exit(1)
	Fatal(v ...interface{})
	// Equivalent to Error() followed by an os.Exit(1), but with additional
	// string interpolation.
	Fatalf(format string, v ...interface{})
}

// Part of handling truly dire situations.  These methods construct a message
// to be provided in a stack trace, allowing the program to attempt
// recovery; unlike FatalWriter which immediately exits.
type PanicWriter interface {
	Panicf(format string, v ...interface{})
}

// Provides a succinct avenue to create new Errors with adequate
// textual information.
type ErrorFactory interface {
	// Creates a new error with the message values interpolated into the
	// format string.
	NewError(format string, v ...interface{}) error
}
