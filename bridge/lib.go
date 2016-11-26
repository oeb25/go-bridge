package bridge

import (
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

type Field struct {
	Name string
	Type string
}

type TargetHeader interface {
	Header() string
}

type TargetFooter interface {
	Footer() string
}

type TargetNamer interface {
	Name(input string, tag reflect.StructTag) string
}

type TargetPtr interface {
	Ptr(to string) string
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

func Format(writer Target, in interface{}) string {
	sum := ""

	if th, exists := writer.(TargetHeader); exists {
		sum = th.Header() + "\n\n"
	}

	_, structs := formatType(writer, reflect.Indirect(reflect.ValueOf(in)).Type())

	var formattedStructs []string

	for i := range structs {
		s := structs[i]
		var sorted []string
		var fields []Field

		for field := range s {
			sorted = append(sorted, field)
		}

		sort.Strings(sorted)

		for i := range sorted {
			field := sorted[i]

			fields = append(fields, Field{
				Name: field,
				Type: s[field],
			})
		}

		this := writer.Struct(i, fields)

		formattedStructs = append(formattedStructs, this)
	}

	sort.Strings(formattedStructs)
	sum = sum + strings.Join(formattedStructs, "\n\n")

	if tf, exists := writer.(TargetFooter); exists {
		sum = sum + tf.Footer()
	}

	return sum + "\n"
}

func formatType(o Target, v reflect.Type) (out string, deps map[string]map[string]string) {
	publicRegex, err := regexp.Compile("^[A-Z]")
	if err != nil {
		log.Fatal(err)
	}

	deps = make(map[string]map[string]string)

	switch v.Kind() {
	case reflect.Struct:
		fields := make(map[string]string)

		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)

			if f.Anonymous {
				continue
			}

			name := f.Name

			// Check if name is exported
			if !publicRegex.Match([]byte(name)) {
				continue
			}

			if n, exists := o.(TargetNamer); exists {
				name = n.Name(name, f.Tag)
			}

			field, dee := formatType(o, f.Type)
			fields[name] = field

			for i := range dee {
				deps[i] = dee[i]
			}
		}

		deps[v.Name()] = fields

		out = v.Name()

	case reflect.Map:
		from, d1 := formatType(o, v.Key())
		for i := range d1 {
			deps[i] = d1[i]
		}
		to, d2 := formatType(o, v.Elem())
		for i := range d2 {
			deps[i] = d2[i]
		}

		out = o.Map(from, to)

	case reflect.Slice:
		out, deps = formatType(o, v.Elem())
		out = o.Array(out)

	case reflect.Ptr:
		n, exists := o.(TargetPtr)

		if !exists {
			// TODO: Add name of target
			log.Fatal("Target does not support pointers, add 'func Ptr(to string) string'")
			return
		}
		out, deps = formatType(o, v.Elem())
		out = n.Ptr(out)

	case reflect.Bool:
		out = o.Bool()
	case reflect.String:
		out = o.String()
	case reflect.Int:
		out = o.Int()
	case reflect.Int16:
		out = o.Int16()
	case reflect.Int32:
		out = o.Int32()
	case reflect.Int64:
		out = o.Int64()
	case reflect.Float32:
		out = o.Float32()
	case reflect.Float64:
		out = o.Float64()
	default:
		log.Fatal("UNHANDELD TYPE:", v.Kind())
		out = v.Name()
	}

	return
}

func dodo(o Target, a interface{}) (name string, structs map[string]map[string]string) {
	v := reflect.Indirect(reflect.ValueOf(a))

	return formatType(o, v.Type())
}