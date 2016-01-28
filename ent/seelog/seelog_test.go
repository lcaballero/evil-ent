package seelog

import (
	"bytes"
	"io"
	"os"
	"testing"

	"fmt"
	"reflect"

	"github.com/lcaballero/evil-ent/ent"
	"github.com/lcaballero/evil-ent/ent/lookup"
	. "github.com/smartystreets/goconvey/convey"
)

func CaptureStdout(print func()) (string, error) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}
	os.Stdout = w
	outC := make(chan string, 0)
	print()

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	return out, nil
}

func NewLogWriter() (ent.LogWriter, error) {
	return NewSeeLogWriter("seelog.xml")
}

func NewLogWriterMocked() (ent.LogWriter, *MockLogger, error) {
	logger := NewMockLogger()
	w := &SeeLogWriter{
		Logger: logger,
		Exit:   func() {},
	}
	return w, logger, nil
}

func TestSeelogConsole(t *testing.T) {

	Convey("Handles creating new logger from config file", t, func() {
		old := LoggerFromConfigAsFile
		defer func() {
			LoggerFromConfigAsFile = old
		}()

		LoggerFromConfigAsFile = badLoggerFromConfigAsFile

		log, err := NewLogWriter()
		So(log, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "badLoggerFromConfigAsFile")
	})

	Convey("After calling Close closes underlying logger should report as closed", t, func() {
		log, mock, err := NewLogWriterMocked()
		So(err, ShouldBeNil)
		So(mock.IsClosed, ShouldBeFalse)
		log.Close()
		So(mock.IsClosed, ShouldBeTrue)
		So(log.Closed(), ShouldBeTrue)
	})

	Convey("Close closes underlying logger", t, func() {
		log, mock, err := NewLogWriterMocked()
		So(err, ShouldBeNil)
		So(mock.IsClosed, ShouldBeFalse)
		log.Close()
		So(mock.IsClosed, ShouldBeTrue)
	})

	Convey("Close sets logger to Closed", t, func() {
		log, mock, err := NewLogWriterMocked()
		defer func() {
			recover()
			log.Flush()
			HasRecord(mock, "Error", "Hello panic\n")
			So(mock.IsFlushed, ShouldBeTrue)
		}()
		So(err, ShouldBeNil)
		log.Panicf("Hello %s", "panic")
	})

	Convey("NewError should produce expected error", t, func() {
		log, err := NewLogWriter()
		defer func() {
			e := recover()
			err, ok := e.(error)
			So(ok, ShouldBeTrue)
			So(err.Error(), ShouldEqual, "Hello panic")
		}()
		So(err, ShouldBeNil)
		log.Panicf("Hello %s", "panic")
	})

	Convey("NewError should produce expected error", t, func() {
		log, err := NewLogWriter()
		So(err, ShouldBeNil)

		r := log.NewError("Hello, %s", "new error")

		So(r, ShouldNotBeNil)
		So(r.Error(), ShouldEqual, "Hello, new error")
	})

	Convey("Error produces expected output", t, func() {
		log, mock, _ := NewLogWriterMocked()
		mapping := lookup.NewLookup(log).MethodMap()
		set := []string{"Debugf", "Errorf", "Fatalf", "Infof", "Tracef", "Warnf"}

		for _, name := range set {
			face := mapping[name]
			param := fmt.Sprintf("Hello %s", name)
			params := []reflect.Value{
				reflect.ValueOf("Hello %s"),
				reflect.ValueOf(name),
			}
			face.Value.Call(params)
			if name == "Fatalf" {
				HasRecord(mock, "Errorf", param)
			} else {
				HasRecord(mock, name, param)
			}
		}
	})

	Convey("Error produces expected output", t, func() {
		log, mock, _ := NewLogWriterMocked()
		mapping := lookup.NewLookup(log).MethodMap()
		set := []string{"Debug", "Error", "Fatal", "Info", "Trace", "Warn"}

		for _, name := range set {
			face := mapping[name]
			param := fmt.Sprintf("Hello %s\n", name)
			params := []reflect.Value{
				reflect.ValueOf("Hello"),
				reflect.ValueOf(name),
			}
			face.Value.Call(params)
			if name == "Fatal" {
				HasRecord(mock, "Error", param)
			} else {
				HasRecord(mock, name, param)
			}
		}
	})

	Convey("Print produces expected output", t, func() {
		log, err := NewLogWriter()
		So(err, ShouldBeNil)

		pr := func() {
			log.Print("Hello")
		}

		s, err := CaptureStdout(pr)
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "Hello")
	})

	Convey("Println produces expected output", t, func() {
		log, err := NewLogWriter()
		So(err, ShouldBeNil)

		pr := func() {
			log.Println("1", "2")
		}

		s, err := CaptureStdout(pr)
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "1 2\n")
	})

	Convey("Println produces expected output", t, func() {
		log, err := NewLogWriter()
		So(err, ShouldBeNil)

		pr := func() {
			log.Printf("answer is %d", 42)
		}

		s, err := CaptureStdout(pr)
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "answer is 42")
	})
}
