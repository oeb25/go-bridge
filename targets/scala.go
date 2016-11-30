package targets

import (
	"reflect"

	"regexp"

	"strings"

	"github.com/oeb25/go-bridge"
)

type Scala struct{}

func (t Scala) Format(in ...interface{}) (string, error) {
	g := bridge.NewBridge(Scala{})
	err := g.FormatMany(in)
	if err != nil {
		return "", err
	}
	return g.Concat(), nil
}

/*
func (t Scala) FormatTo(in interface{}, path string) error {
	types, err := t.Format(in)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(types), 0700)
}
*/

func (t Scala) Name(input string, tags reflect.StructTag) string {
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

func (t Scala) Map(from, to string) string {
	return "Map[" + from + ", " + to + "]"
}

func (t Scala) Array(of string) string {
	return "List[" + of + "]"
}

func (t Scala) Struct(name string, fields []bridge.Field) (out string) {
	out = "private type " + name + " = {\n"
	for n := range fields {
		f := fields[n]
		out = out + "  val " + f.Name + ": " + f.Type + "\n"
	}

	out = out + "}"

	return
}

func (t Scala) Ptr(to string) string {
	return "Box<" + to + ">"
}

func (t Scala) Int() string   { return "Int" }
func (t Scala) Int8() string  { return "Byte" }
func (t Scala) Int16() string { return "Short" }
func (t Scala) Int32() string { return "Int" }
func (t Scala) Int64() string { return "Long" }

func (t Scala) Float32() string { return "Float" }
func (t Scala) Float64() string { return "Double" }

func (t Scala) String() string { return "String" }

func (t Scala) Bool() string { return "Boolean" }
