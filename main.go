package main

import (
	"log"

	"github.com/oeb25/go-bridge/targets"
)

type User struct {
	ID      int      `json:"id"`
	Friends []Friend `json:"friends"`
}

type Friend struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"firstname"`
}

func main() {
	err := targets.TypeScript{}.
		FormatTo(User{}, "./types.ts")
	if err != nil {
		log.Fatal(err)
	}
}
