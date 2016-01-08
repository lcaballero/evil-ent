package main

import (
	"fmt"

	"github.com/lcaballero/evil-ent/ent/lookup"
	"github.com/lcaballero/evil-ent/ent/seelog"
)

func main() {
	Dump()
}

func Dump() {
	p := lookup.NewLookup(&seelog.SeeLogWriter{})
	methods := p.Methods()

	for _, m := range methods {
		fmt.Println(m.Name())
	}
}
