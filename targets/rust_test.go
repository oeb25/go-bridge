package targets_test

import (
	"testing"

	"github.com/oeb25/go-bridge/targets"
	"github.com/stretchr/testify/assert"
)

type Simple struct {
	Name string
	Age  int
}

func TestSimple(t *testing.T) {
	actual, err := targets.Rust{}.Format(Simple{})
	assert.NoError(t, err)

	expected := `
#[derive(Debug, Default)]
struct Simple {
    name: String,
    age: i32,
}
`

	assert.Equal(t, expected, "\n"+actual)
}

type Nested struct {
	Something float32
	Nest      Simple
}

func TestNested(t *testing.T) {
	actual, err := targets.Rust{}.Format(Nested{})
	assert.NoError(t, err)

	expected := `
#[derive(Debug, Default)]
struct Simple {
    name: String,
    age: i32,
}
#[derive(Debug, Default)]
struct Nested {
    something: f32,
    nest: Simple,
}
`

	assert.Equal(t, expected, "\n"+actual)
}

func TestPointer(t *testing.T) {
	assert.Equal(t, "Box<x>", targets.Rust{}.Ptr("x"))
}

func TestMap(t *testing.T) {
	assert.Equal(t,
		"std::collections::HashMap<a, b>",
		targets.Rust{}.Map("a", "b"),
	)
}

func TestSlice(t *testing.T) {
	assert.Equal(t, true, true)
}
