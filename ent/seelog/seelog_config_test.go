package seelog

import (
	"testing"

	"encoding/xml"

	. "github.com/smartystreets/goconvey/convey"
)

const seelog_xml = `
<?xml version="1.0"?>
<seelog type="asynctimer" asyncinterval="1000000" minlevel="info">
  <outputs formatid="all">
    <filter levels="trace,debug,info,warn">
      <rollingfile type="size" filename="/var/log/persistent/evilent.log" maxsize="20000000" maxrolls="5"/>
    </filter>
    <filter levels="error,critical" formatid="fmterror">
      <rollingfile type="size" filename="/var/log/persistent/evilent.log" maxsize="20000000" maxrolls="5"/>
    </filter>
  </outputs>
  <formats>
    <format id="all" format="%Date %Time [%LEVEL] - %Msg%n"/>
    <format id="fmterror" format="%Date %Time [%LEVEL] [%FuncShort @ %File.%Line] - %Msg%n"/>
  </formats>
</seelog>
`

func ReadSeelogConf() *SeelogConf {
	v := SeelogConf{}
	err := xml.Unmarshal([]byte(seelog_xml), &v)
	So(err, ShouldBeNil)
	return &v
}

func TestSeelogConfig(t *testing.T) {

	Convey("Outputs[0].Filters[0] should have levels", t, func() {
		v := ReadSeelogConf()
		So(v.Outputs[0].Filters[0].Levels, ShouldEqual, "trace,debug,info,warn")
	})

	Convey("Each format should have an id and a format", t, func() {
		v := ReadSeelogConf()
		format := v.Formats[0]
		So(format.Id, ShouldEqual, "all")

		So(format.Format, ShouldContainSubstring, "Date")
		So(format.Format, ShouldContainSubstring, "Time")
		So(format.Format, ShouldContainSubstring, "LEVEL")

		format = v.Formats[1]
		So(format.Id, ShouldEqual, "fmterror")

		So(format.Format, ShouldContainSubstring, "Date")
		So(format.Format, ShouldContainSubstring, "Time")
		So(format.Format, ShouldContainSubstring, "FuncShort")
	})

	Convey("Formats should have length 2", t, func() {
		v := ReadSeelogConf()
		So(v.Formats, ShouldNotBeNil)
		So(len(v.Formats), ShouldEqual, 2)
	})

	Convey("Outputs[0] should have 2 filter instances", t, func() {
		v := ReadSeelogConf()
		So(len(v.Outputs[0].Filters), ShouldEqual, 2)
	})

	Convey("Outputs should have length 1 and formatid of 'all'", t, func() {
		v := ReadSeelogConf()
		So(v.Outputs, ShouldNotBeNil)
		So(len(v.Outputs), ShouldEqual, 1)
		So(v.Outputs[0].FormatId, ShouldEqual, "all")
	})

	Convey("Should find root element", t, func() {
		v := ReadSeelogConf()
		So(v.XMLName.Local, ShouldEqual, "seelog")
	})

	Convey("Seelog type should be 'asynctimer'", t, func() {
		v := ReadSeelogConf()
		So(v.Type, ShouldEqual, "asynctimer")
	})

	Convey("Seelog asyncinterval should be '1000000'", t, func() {
		v := ReadSeelogConf()
		So(v.AsyncInterval, ShouldEqual, 1000000)
	})
}
