package seelog

import (
	"fmt"
	"os"

	"github.com/lcaballero/evil-ent/ent"
	log "github.schq.secious.com/Logrhythm/Godeps/_workspace/src/github.com/cihub/seelog"
	"errors"
)

// Can be overwritten for purposes of testing.
var LoggerFromConfigAsFile = log.LoggerFromConfigAsFile
var badLoggerFromConfigAsFile = func(fileName string) (log.LoggerInterface, error) {
	return nil, errors.New("badLoggerFromConfigAsFile")
}

// SeeLogWriter implements the log.LogWriter interface and all of it's
// delegates.
type SeeLogWriter struct {
	Logger      DelgateInterface
	RecentError error
	Exit        ent.ExitStrategy
}

// Creates a new SeeLogWriter, or produces an error if one should occur during
// initial creation.
func NewSeeLogWriter(configxml string) (*SeeLogWriter, error) {
	logger, err := LoggerFromConfigAsFile(configxml)

	if err != nil {
		log.Warn("Failed to load config", err)
		return nil, err
	} else {
		log.ReplaceLogger(logger)
		logger.Info("replaced logger")
	}

	w := &SeeLogWriter{
		Logger: logger,
		Exit: func() {
			os.Exit(1)
		},
	}

	return w, nil
}

// Error meets LogWriter interface for writing values at level = Error
func (w *SeeLogWriter) Error(v ...interface{}) {
	w.captureError(w.Logger.Error(v...))
}

// Logs the interpolated value with level = Error
func (w *SeeLogWriter) Errorf(format string, v ...interface{}) {
	w.captureError(w.Logger.Errorf(format, v...))
}

// Logs the values provided with level = Debug
func (w *SeeLogWriter) Debug(v ...interface{}) {
	w.Logger.Debug(v...)
}

// Logs the interpolated value with level = Debug
func (w *SeeLogWriter) Debugf(format string, v ...interface{}) {
	w.Logger.Debugf(format, v...)
}

// Logs the values provided with level = Warn
func (w *SeeLogWriter) Warn(v ...interface{}) {
	w.captureError(w.Logger.Warn(v...))
}

// Logs the interpolated value with level = Warn
func (w *SeeLogWriter) Warnf(format string, v ...interface{}) {
	w.captureError(w.Logger.Warnf(format, v...))
}

// Logs the values provided with level = Info
func (w *SeeLogWriter) Info(v ...interface{}) {
	w.Logger.Info(v...)
}

// Logs the interpolated value with level = Info
func (w *SeeLogWriter) Infof(format string, v ...interface{}) {
	w.Logger.Infof(format, v...)
}

// Logs the values provided with level = Trace
func (w *SeeLogWriter) Trace(v ...interface{}) {
	w.Logger.Trace(v...)
}

// Trace (LogWriter interface) interpolates and writes values at level = Trace
func (w *SeeLogWriter) Tracef(format string, v ...interface{}) {
	w.Logger.Tracef(format, v...)
}

// Equivalent to Error() followed by an os.Exit(1)
func (w *SeeLogWriter) Fatal(v ...interface{}) {
	w.captureError(w.Logger.Error(v...))
	w.Exit()
}

// Creates a new error with the message values interpolated into the
// format string.
func (w *SeeLogWriter) NewError(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}

// Equivalent to Error() followed by an os.Exit(1), but with additional
// string interpolation.
func (w *SeeLogWriter) Fatalf(format string, v ...interface{}) {
	w.captureError(w.Logger.Errorf(format, v...))
	w.Exit()
}

// Panicf creates a new error with the given format and values then
// calls panic with that error.
func (w *SeeLogWriter) Panicf(format string, v ...interface{}) {
	w.captureError(w.NewError(format, v...))
	panic(w.RecentError)
}

// Captures any errors that were reported by the underlying logger when
// calling normal logging functions.
func (w *SeeLogWriter) captureError(err error) {
	if err != nil {
		w.RecentError = err
	}
}

// Flush writes any held error to the log and then calls flush on underlying
// logger.
func (w *SeeLogWriter) Flush() {
	w.Error(w.RecentError)
	w.Logger.Flush()
}

// Close closes this logger and any underlying logger.
func (w *SeeLogWriter) Close() error {
	w.Logger.Close()
	return nil
}

// Closed returns true if the underlying logger has been closed and false
// otherwise.
func (w *SeeLogWriter) Closed() bool {
	return w.Logger.Closed()
}

// Print outputs to Stdout.
func (w *SeeLogWriter) Print(v ...interface{}) {
	fmt.Print(v...)
}

// Println outputs to Stdout.
func (w *SeeLogWriter) Println(v ...interface{}) {
	fmt.Println(v...)
}

// Printf outputs to Stdout.
func (w *SeeLogWriter) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
