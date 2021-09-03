package main

import (
	"fmt"
	"reflect"
	"time"

	. "github.com/dave/jennifer/jen"
)

func main() {
	t := reflect.TypeOf(time.Time{})
	f := NewFile("chrono")
	exclude := map[string]bool{
		"MarshalJSON": true,
		"String":      true,
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		mt := m.Type

		if exclude[m.Name] {
			continue
		}

		params := make([]Code, mt.NumIn()-1)
		paramNames := make([]Code, mt.NumIn()-1)

		for j := 1; j < mt.NumIn(); j++ {
			in := mt.In(j)
			p := Id(fmt.Sprintf("p%d", j))
			paramNames[j-1] = Id(fmt.Sprintf("p%d", j))

			switch in.Kind() {
			case reflect.Ptr:
				p = p.Op("*")
				in = in.Elem()
			case reflect.Slice:
				p = p.Op("[]")
				in = in.Elem()
			}

			params[j-1] = p.Qual(in.PkgPath(), in.Name())
		}

		returns := make([]Code, mt.NumOut())

		for j := 0; j < mt.NumOut(); j++ {
			out := mt.Out(j)
			returns[j] = Id(out.String())
		}

		f.Func().Params(Id("t").Op("*").Id("Time")).
			Id(m.Name).Params(params...).Call(returns...).Block(
			Return(Id("t").Dot("T").Call().Dot(m.Name).Call(paramNames...)),
		)
	}

	if err := f.Save("stubs.gen.go"); err != nil {
		panic(err)
	}
}
