package main

import (
	"log"

	"gitlab.com/oeb25/go-bridge/targets"
)

type types struct {
	Order      Order
	TypeScript targets.TypeScript
}

func main() {
	err := targets.Rust{}.FormatTo(types{}, "./types.rs")
	if err != nil {
		log.Fatal(err)
	}
}
