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
	LastName  string `json:"lastname"`
}

func main() {
	err := targets.C{}.
		FormatTo(User{}, "./types.c")
	if err != nil {
		log.Fatal(err)
	}
}
