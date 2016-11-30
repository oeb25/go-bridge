package main

import (
	"fmt"

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
	g, err := targets.TypeScript{}.Format(User{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g)
}
