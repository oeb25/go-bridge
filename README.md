# go-bridge

[![GoDoc](https://godoc.org/github.com/oeb25/go-bridge?status.svg)](https://godoc.org/github.com/oeb25/go-bridge)
[![Build Status](https://travis-ci.org/oeb25/go-bridge.svg?branch=master)](https://travis-ci.org/oeb25/go-bridge)

Convert your Go structs to other languages, including TypeScript / Flow, Elm and Rust among others!

```go
type User struct {
	ID      int      `json:"id"`
	Friends []Friend `json:"friends"`
}

type Friend struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
```

in TypeScript becomes...

```typescript
export interface User {
	friends: Friend[],
	id: number,
}

export interface Friend {
	firstname: string,
	lastname: string,
}
```

## Usage

First get go-bridge
```bash
go get github.com/oeb25/go-bridge
```

Then import and use it in your code

```go
import (
	"fmt"
	"github.com/oeb25/go-bridge/targets"
)

func main() {
	result, _ := targets.TypeScript{}.Format(MyStruct{})
	fmt.Println(result)
}
```

## Officially supported targets

- TypeScript / Flow
- Elm _(partial)_
- Rust _(partial)_
- C _(**very** much work in progress)_
