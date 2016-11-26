# go-bridge

Compile your Go structs to other languages!

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
import "github.com/oeb25/go-bridge/targets"

func main() {
	targets.TypeScript{}.FormatTo(MyStruct{}, "./types.ts")
}
```

## Officially supported targets

- TypeScript / Flow
- Rust _(partial)_ 
