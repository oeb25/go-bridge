package bridge

import "reflect"

type Header interface {
	Header() string
}

type Footer interface {
	Footer() string
}

type Namer interface {
	Name(input string, tag reflect.StructTag) string
}

type Any interface {
	Any() string
}

type Ptr interface {
	Ptr(to string) string
}

type Func interface {
	Func(name string, args []string, out string) string
}

type FuncMultipleReturn interface {
	Func(name string, args []string, out []string) string
}

type Interface interface {
	Interface(name string, methods []string) string
}

type Target interface {
	Struct(name string, fields []Field) string
	Map(from string, to string) string
	Array(of string) string

	Int() string
	Int8() string
	Int16() string
	Int32() string
	Int64() string

	Float32() string
	Float64() string

	String() string
	Bool() string
}
