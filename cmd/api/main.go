package main

import (
	"github.com/juanegido/hexapi/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		panic(err)
	}
}
