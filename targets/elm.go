package targets

import (
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/oeb25/go-bridge"
)

type Elm struct{}

func (t Elm) Format(in interface{}) (string, error) {
	return bridge.Format(t, in)
}

func (t Elm) FormatTo(in interface{}, path string) error {
	types, err := t.Format(in)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(types), 0700)
}

func (t Elm) Header() string {
	return "import Dict exposing ( Dict )"
}

func (t Elm) Name(input string, tags reflect.StructTag) string {
	j := tags.Get("json")
	if j == "" {
		return input
	}

	return strings.Split(j, ",")[0]
}

func (t Elm) Map(from, to string) string {
	return "Dict " + from + " " + to
}

func (t Elm) Array(of string) string {
	return "List " + of
}

func (t Elm) Struct(name string, fields []bridge.Field) (out string) {
	out = "type alias " + name + " =\n"

	var fs []string

	for n := range fields {
		f := fields[n]
		fs = append(fs, f.Name+" : "+f.Type)
	}

	out = out + "  { " + strings.Join(fs, "\n  , ")

	out = out + "\n  }"

	return
}

func (t Elm) Ptr(to string) string {
	return to
}

func (t Elm) Int() string   { return "Int" }
func (t Elm) Int8() string  { return "Int" }
func (t Elm) Int16() string { return "Int" }
func (t Elm) Int32() string { return "Int" }
func (t Elm) Int64() string { return "Int" }

func (t Elm) Float32() string { return "Float" }
func (t Elm) Float64() string { return "Float" }

func (t Elm) String() string { return "String" }

func (t Elm) Bool() string { return "Bool" }
