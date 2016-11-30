// Package bridge converts your Go structs to other languages, including
// TypeScript / Flow, Elm and Rust among others!
package bridge

import (
	"errors"
	"reflect"
)

type Field struct {
	Name string
	Type string
}

type Method struct {
	Name string
	Args []string
	Outs []string
}

type definition struct {
	Type   reflect.Type
	Deps   map[reflect.Type]string
	Name   string
	Parsed string
}

type Bridge struct {
	Target      Target
	definitions []definition
	flag        map[reflect.Type]bool
}

func NewBridge(target Target) Bridge {
	return Bridge{
		Target: target,
		flag:   make(map[reflect.Type]bool),
	}
}

func (g *Bridge) getDependencies(t reflect.Type) (deps []reflect.Type) {
	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			deps = append(deps, t.Field(i).Type)
		}
	}

	return
}

func (g *Bridge) addDefinition(t reflect.Type) *definition {
	for i := range g.definitions {
		def := g.definitions[i]
		if def.Type == t {
			return &def
		}
	}

	def := definition{
		Type: t,
		Deps: make(map[reflect.Type]string),
	}

	g.definitions = append(g.definitions, def)

	return &def
}

func (g *Bridge) findType(t reflect.Type) *definition {
	for i := range g.definitions {
		def := g.definitions[i]
		if def.Type == t {
			return &def
		}
	}

	return nil
}

func (g *Bridge) parse(d *definition) (err error) {
	targetName := reflect.TypeOf(g.Target).Name()

	t := d.Type
	switch t.Kind() {
	case reflect.Struct:
		fields := make([]Field, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			dd := g.findType(t.Field(i).Type)
			name := t.Field(i).Name
			if o, exists := g.Target.(Namer); exists {
				name = o.Name(name, t.Field(i).Tag)
			}
			fields[i] = Field{
				Name: name,
				Type: dd.Name,
			}
		}
		d.Name = t.Name()
		d.Parsed = g.Target.Struct(t.Name(), fields)

	case reflect.Map:
		dKey := g.findType(t.Key())
		key := dKey.Name
		switch dKey.Type.Kind() {
		case reflect.Interface, reflect.Struct:
			key = dKey.Type.Name()
		}

		dElem := g.findType(t.Elem())
		elem := dElem.Name
		switch dElem.Type.Kind() {
		case reflect.Interface, reflect.Struct:
			elem = dElem.Type.Name()
		}

		d.Name = g.Target.Map(key, elem)

	case reflect.Slice:
		dd := g.findType(t.Elem())
		name := dd.Name
		switch dd.Type.Kind() {
		case reflect.Interface, reflect.Struct:
			name = dd.Type.Name()
		}
		d.Name = g.Target.Array(name)

	case reflect.Ptr:
		var o Ptr
		var exists bool
		if o, exists = g.Target.(Ptr); !exists {
			err = errors.New("Target " + targetName + " does not have function Ptr")
			return
		}
		d.Name = o.Ptr(g.findType(t.Elem()).Name)

	case reflect.Interface:
		d.Name = t.Name()
		if t.NumMethod() == 0 && d.Name == "" {
			if o, e := g.Target.(Any); e {
				d.Name = o.Any()
				return
			}
			err = errors.New("Target " + reflect.TypeOf(g.Target).Name() +
				" does not support any, needed by an empty unnamed interface. " +
				"To fix this add func Any to the target")
			return
		}

		methods := make([]string, t.NumMethod())
		for i := 0; i < t.NumMethod(); i++ {
			method := t.Method(i)
			dd := g.findType(method.Type)
			methods[i] = method.Name + dd.Name
		}
		var o Interface
		var exists bool
		if o, exists = g.Target.(Interface); !exists {
			err = errors.New("Target " + reflect.TypeOf(g.Target).Name() + " does not support interfaces")
			return
		}

		d.Parsed = o.Interface(d.Name, methods)

	case reflect.Func:
		args := make([]string, t.NumIn())
		outs := make([]string, t.NumOut())

		for i := 0; i < t.NumIn(); i++ {
			k := g.findType(t.In(i))
			if k.Type.Kind() == reflect.Interface {
				args[i] = k.Type.Name()
			} else {
				args[i] = k.Name
			}
		}
		if t.NumOut() > 1 {
			err = errors.New("Target " + reflect.TypeOf(g.Target).Name() + " multiple return arguments")
			return

		}
		for i := 0; i < t.NumOut(); i++ {
			k := g.findType(t.Out(i))
			if k.Type.Kind() == reflect.Interface {
				outs[i] = k.Type.Name()
			} else {
				outs[i] = k.Name
			}
		}

		var o Func
		var exists bool
		if o, exists = g.Target.(Func); !exists {
			err = errors.New("Target " + reflect.TypeOf(g.Target).Name() + " does not have function Func")
			return
		}

		d.Name = o.Func(t.Name(), args, outs[0])

	case reflect.Bool:
		d.Name = g.Target.Bool()
	case reflect.String:
		d.Name = g.Target.String()
	case reflect.Int:
		d.Name = g.Target.Int()
	case reflect.Int16:
		d.Name = g.Target.Int16()
	case reflect.Int32:
		d.Name = g.Target.Int32()
	case reflect.Int64:
		d.Name = g.Target.Int64()
	case reflect.Float32:
		d.Name = g.Target.Float32()
	case reflect.Float64:
		d.Name = g.Target.Float64()
	default:
		err = errors.New("UNHANDELD TYPE: " + t.Kind().String())
		return
	}

	return
}

func (g *Bridge) resolve(t reflect.Type) {
	if _, exists := g.flag[t]; exists {
		return
	}
	g.flag[t] = true
	defer g.addDefinition(t)

	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			g.resolve(t.Field(i).Type)
		}
	case reflect.Interface:
		for i := 0; i < t.NumMethod(); i++ {
			g.resolve(t.Method(i).Type)
		}
	case reflect.Map:
		g.resolve(t.Key())
		g.resolve(t.Elem())
	case reflect.Ptr:
		g.resolve(t.Elem())
	case reflect.Slice:
		g.resolve(t.Elem())
	case reflect.Func:
		for i := 0; i < t.NumIn(); i++ {
			g.resolve(t.In(i))
		}
		for i := 0; i < t.NumOut(); i++ {
			g.resolve(t.Out(i))
		}
	}
}

func (g *Bridge) Format(ts ...interface{}) error {
	return g.FormatMany(ts)
}

func (g *Bridge) FormatMany(ts []interface{}) error {
	for i := range ts {
		g.resolve(reflect.TypeOf(ts[i]))
	}
	for i := range g.definitions {
		if err := g.parse(&g.definitions[i]); err != nil {
			return err
		}
	}

	return nil
}

func (g *Bridge) Concat() (out string) {
	if o, e := g.Target.(Header); e {
		out += o.Header() + "\n"
	}
	for i := range g.definitions {
		if g.definitions[i].Parsed != "" {
			out += g.definitions[i].Parsed + "\n"
		}
	}
	if o, e := g.Target.(Footer); e {
		out += o.Footer() + "\n"
	}

	return
}
