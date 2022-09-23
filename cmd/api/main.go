package main

import (
	"log"

	"github.com/juanegido/hexapi/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
