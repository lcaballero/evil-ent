package seelog

import "encoding/xml"

type SeelogConf struct {
	XMLName       xml.Name  `xml:"seelog"`
	Type          string    `xml:"type,attr"`
	AsyncInterval int       `xml:"asyncinterval,attr"`
	MinLevel      string    `xml:"minlevel,attr"`
	Outputs       []Outputs `xml:"outputs"`
	Formats       []Format  `xml:"formats>format"`
}
type Outputs struct {
	XNLName  xml.Name `xml:"outputs"`
	Filters  []Filter `xml:"filter"`
	FormatId string   `xml:"formatid,attr"`
}

type Filter struct {
	Levels      string `xml:"levels,attr"`
	RollingFile RollingFile
}
type RollingFile struct {
	Type     string `xml:"size,attr"`
	Filename string `xml:"filename,attr"`
	MaxSize  int    `xml:"maxsize,attr"`
	MaxRolls int    `xml:"maxrolls,attr"`
}
type Format struct {
	Id     string `xml:"id,attr"`
	Format string `xml:"format,attr"`
}
