package targets

import (
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/oeb25/go-bridge/bridge"
)

type TypeScript struct {
	Flow      bool
	Immutable bool
}

func (t TypeScript) Header() string {
	var header []string

	if t.Flow {
		header = append(header, "// @flow")
	}
	if t.Immutable {
		header = append(header, "import { Map, List } from 'immutable'")
	}

	header = append(header, "// WARNING: THIS FILE IS AUTOGENERATED! ANY CHANGES WILL BE OVERWRITTEN WITHOUT WARNING")

	return strings.Join(header, "\n\n")
}

func (t TypeScript) Format(in interface{}) string {
	return bridge.Format(t, in)
}

func (t TypeScript) FormatTo(in interface{}, path string) error {
	types := t.Format(in)
	return ioutil.WriteFile(path, []byte(types), 0700)
}

func (t TypeScript) Name(input string, tags reflect.StructTag) string {
	j := tags.Get("json")
	if j == "" {
		return input
	}

	return strings.Split(j, ",")[0]
}

func (t TypeScript) Map(from, to string) string {
	if t.Immutable {
		return "Map<" + from + ", " + to + ">"
	} else {
		return "{ [key: " + from + "]: " + to + " }"
	}
}

func (t TypeScript) Array(of string) string {
	if t.Immutable {
		return "List<" + of + ">"
	} else {
		return of + "[]"
	}
}

func (t TypeScript) Struct(name string, fields []bridge.Field) (out string) {
	out = "export interface " + name + " {\n"
	for n := range fields {
		f := fields[n]
		out = out + "\t" + f.Name + ": " + f.Type + ",\n"
	}

	out = out + "}"

	return
}

func (t TypeScript) Ptr(to string) string {
	return to
}

func (t TypeScript) Int() string   { return "number" }
func (t TypeScript) Int8() string  { return "number" }
func (t TypeScript) Int16() string { return "number" }
func (t TypeScript) Int32() string { return "number" }
func (t TypeScript) Int64() string { return "number" }

func (t TypeScript) Float32() string { return "number" }
func (t TypeScript) Float64() string { return "number" }

func (t TypeScript) String() string { return "string" }

func (t TypeScript) Bool() string { return "boolean" }
