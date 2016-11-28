package targets

import (
	"io/ioutil"
	"reflect"

	"regexp"

	"strings"

	"log"

	"fmt"

	"github.com/oeb25/go-bridge"
)

type C struct {
	BoolType    string
	StringType  string
	HashMapType string
}

func (t C) Format(in interface{}) (string, error) {
	return bridge.Format(t, in)
}

func (t C) FormatTo(in interface{}, path string) error {
	types, err := t.Format(in)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(types), 0700)
}

func (t C) Name(input string, tags reflect.StructTag) string {
	r, _ := regexp.Compile("([A-Z]+[a-z0-9]+)")
	out := r.FindAllString(input, -1)

	if len(out) == 0 {
		return strings.ToLower(input)
	}

	for i := range out {
		out[i] = strings.ToLower(out[i])
	}

	return strings.Join(out, "_")
}

func (t C) Map(from, to string) string {
	log.Fatal("Conversion to hashmaps in C is currently not supported")
	return ""
}

func (t C) Array(of string) string {
	fmt.Println("WARNING: Arrays in C are currently just pointers to the type, be careful")
	return of + " *"
}

func (t C) Struct(name string, fields []bridge.Field) (out string) {
	out = "struct " + name + " {\n"
	for n := range fields {
		f := fields[n]
		out = out + "\t" + f.Type + " " + f.Name + ";\n"
	}

	out = out + "};"

	return
}

func (t C) Ptr(to string) string {
	return to + " *"
}

func (t C) Int() string   { return "int" }
func (t C) Int8() string  { return "signed char" }
func (t C) Int16() string { return "int" }
func (t C) Int32() string { return "long int" }
func (t C) Int64() string { return "long long int" }

func (t C) Float32() string { return "float" }
func (t C) Float64() string { return "double" }

func (t C) String() string {
	if t.StringType != "" {
		return t.StringType
	}

	return "char *"
}

func (t C) Bool() string {
	if t.BoolType != "" {
		return t.BoolType
	}

	return "signed char"
}
