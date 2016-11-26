package targets

import (
	"io/ioutil"
	"reflect"

	"gitlab.com/oeb25/go-bridge/bridge"
)

type Rust struct{}

func (t Rust) Format(in interface{}) string {
	return bridge.Format(t, in)
}

func (t Rust) FormatTo(in interface{}, path string) error {
	types := t.Format(in)
	return ioutil.WriteFile(path, []byte(types), 0700)
}

func (t Rust) Header() string {
	return "use std::collections::HashMap\n"
}

func (t Rust) Name(input string, tags reflect.StructTag) string {
	return input
}

func (t Rust) Map(from, to string) string {
	return "std::collections::HashMap<" + from + ", " + to + ">"
}

func (t Rust) Array(of string) string {
	return "Vec<" + of + ">"
}

func (t Rust) Struct(name string, fields []bridge.Field) (out string) {
	out = "struct " + name + " {\n"
	for n := range fields {
		f := fields[n]
		out = out + "\t" + f.Name + ": " + f.Type + ",\n"
	}

	out = out + "}\n"

	return
}

func (t Rust) Ptr(to string) string {
	return "Box<" + to + ">"
}

func (t Rust) Int() string   { return "i32" }
func (t Rust) Int8() string  { return "i8" }
func (t Rust) Int16() string { return "i16" }
func (t Rust) Int32() string { return "i32" }
func (t Rust) Int64() string { return "i64" }

func (t Rust) Float32() string { return "f32" }
func (t Rust) Float64() string { return "f64" }

func (t Rust) String() string { return "String" }

func (t Rust) Bool() string { return "bool" }
