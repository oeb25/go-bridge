package bridge_test

import (
	"testing"

	"fmt"

	"strings"

	"github.com/oeb25/go-bridge"
	"github.com/stretchr/testify/assert"
)

type SampleTarget struct{}

func (t SampleTarget) Bool() string               { return "Bool" }
func (t SampleTarget) Int() string                { return "Int" }
func (t SampleTarget) Int8() string               { return "Int8" }
func (t SampleTarget) Int16() string              { return "Int16" }
func (t SampleTarget) Int32() string              { return "Int32" }
func (t SampleTarget) Int64() string              { return "Int64" }
func (t SampleTarget) Float32() string            { return "Float32" }
func (t SampleTarget) Float64() string            { return "Float64" }
func (t SampleTarget) String() string             { return "String" }
func (t SampleTarget) Array(of string) string     { return "[]" + of }
func (t SampleTarget) Map(from, to string) string { return "map[" + from + "]" + to }
func (t SampleTarget) Struct(name string, fields []bridge.Field) (out string) {
	out += "struct " + name + " {"
	for i := range fields {
		out += fields[i].Name + ": " + fields[i].Type + ","
	}
	out += "}"
	return
}
func (t SampleTarget) Interface(name string, methods []string) (out string) {
	return "interface " + name + " { " + strings.Join(methods, ", ") + " }"
}
func (t SampleTarget) Any() string {
	return "ANY"
}

func TestMap(t *testing.T) {
	type A struct {
		A map[string]int
	}

	g := bridge.NewBridge(SampleTarget{})
	assert.NoError(t, g.Format(A{}))
	result := g.Concat()
	fmt.Println(result)
	assert.Equal(t, `
struct A {A: map[String]Int,}
`, "\n"+result)
}

func TestSlice(t *testing.T) {
	type A struct {
		A []int
	}

	g := bridge.NewBridge(SampleTarget{})
	assert.NoError(t, g.Format(A{}))
	result := g.Concat()
	fmt.Println(result)
	assert.Equal(t, `
struct A {A: []Int,}
`, "\n"+result)
}

func TestMultiple(t *testing.T) {
	type A struct {
		A int
	}

	type B struct {
		A A
		B int
	}

	g := bridge.NewBridge(SampleTarget{})
	assert.NoError(t, g.Format(A{}, B{}))
	result := g.Concat()
	assert.Equal(t, `
struct A {A: Int,}
struct B {A: A,B: Int,}
`, "\n"+result)
}

func TestRecursive(t *testing.T) {
	type A struct {
		A []A
	}

	g := bridge.NewBridge(SampleTarget{})
	assert.NoError(t, g.Format(A{}))
	result := g.Concat()
	fmt.Println(result)
	assert.Equal(t, `
struct A {A: []A,}
`, "\n"+result)
}

func TestEmptyInterface(t *testing.T) {
	type API interface{}
	type wrapper struct {
		api API
	}

	g := bridge.NewBridge(SampleTarget{})
	assert.NoError(t, g.Format(wrapper{}))
	result := g.Concat()
	fmt.Println(result)
	assert.Equal(t, `
interface API {  }
struct wrapper {api: API,}
`, "\n"+result)
}

func TestEmptyInterface2(t *testing.T) {
	type wrapper struct {
		api interface{}
	}

	g := bridge.NewBridge(SampleTarget{})
	assert.NoError(t, g.Format(wrapper{}))
	result := g.Concat()
	fmt.Println(result)
	assert.Equal(t, `
struct wrapper {api: ANY,}
`, "\n"+result)
}
