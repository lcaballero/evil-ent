package lookup

import "reflect"

type Lookup struct {
	target interface{}
}

type Facet struct {
	Method  reflect.Method
	Type    reflect.Type
	Value   reflect.Value
	ValueOf reflect.Value
}

func (f *Facet) Name() string {
	names := []string{
		f.Method.Name,
		f.Type.Name(),
	}

	for _, name := range names {
		if name != "" {
			return name
		}
	}
	return ""
}

func NewLookup(a interface{}) *Lookup {
	return &Lookup{
		target: a,
	}
}

func (p *Lookup) Methods() []*Facet {
	ty := reflect.TypeOf(p.target)
	val := reflect.ValueOf(p.target)

	methods := make([]*Facet, ty.NumMethod())

	for i := range methods {
		methods[i] = &Facet{
			Method:  ty.Method(i),
			Type:    ty,
			Value:   val.Method(i),
			ValueOf: val,
		}
	}

	return methods
}

func (p *Lookup) MethodMap() map[string]*Facet {
	methods := p.Methods()
	faces := make(map[string]*Facet)
	for _, f := range methods {
		faces[f.Name()] = f
	}
	return faces
}
