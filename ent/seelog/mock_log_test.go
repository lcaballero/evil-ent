package seelog

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/lcaballero/evil-ent/ent/lookup"
	"fmt"
	"reflect"
)


func HasRecord(log *MockLogger, key, val string) {
	vals := log.Writes[key]

	So(vals, ShouldNotBeNil)

	n := len(vals)
	So(n, ShouldBeGreaterThanOrEqualTo, 1)
	So(vals[n-1], ShouldEqual, val) // Checks the 'last'
}

func TestMockLog(t *testing.T) {

	Convey("Call all of the non-format logging methods.", t, func() {
		log := NewMockLogger()
		mapping  := lookup.NewLookup(log).MethodMap()

		set := []string{ "Debug", "Error", "Fatal", "Info", "Trace", "Warn" }

		for _,name := range set {
			face := mapping[name]
			param := fmt.Sprintf("Hello %s\n", name)
			params := []reflect.Value{
				reflect.ValueOf("Hello"),
				reflect.ValueOf(name),
			}
			face.Value.Call(params)
			HasRecord(log, name, param)
		}
	})

	Convey("Call all of the formatted logging methods.", t, func() {
		log := NewMockLogger()
		mapping := lookup.NewLookup(log).MethodMap()

		set := []string{
			"Debugf", "Errorf", "Fatalf", "Infof", "Tracef", "Warnf",
		}

		for _,name := range set {
			face := mapping[name]
			param := fmt.Sprintf("Hello %s", name)
			params := []reflect.Value{
				reflect.ValueOf("Hello %s"),
				reflect.ValueOf(name),
			}
			face.Value.Call(params)
			HasRecord(log, name, param)
		}
	})


	Convey("MockLogger should record Debugf(...) call.", t, func() {
		log := NewMockLogger()
		log.Debugf("Hello, %s", "World")
		HasRecord(log, "Debugf", "Hello, World")
	})

	Convey("MockLogger should record Errorf(...) call.", t, func() {
		log := NewMockLogger()
		log.Errorf("Hello, %s", "World")
		HasRecord(log, "Errorf", "Hello, World")
	})

	Convey("MockLogger should record Error(...) call.", t, func() {
		log := NewMockLogger()
		log.Error("Hello", "World")
		HasRecord(log, "Error", "Hello World\n")
	})

	Convey("MockLogger should instantiate without incident.", t, func() {
		log := NewMockLogger()
		So(log, ShouldNotBeNil)
		So(log.Writes, ShouldNotBeNil)
	})
}
