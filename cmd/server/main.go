package main

import (
	"github.com/storyscript/fakeruntime/http"
)

func main() {
	server := http.Server{}

	if err := server.Start(); err != nil {
		panic(err)
	}
}
