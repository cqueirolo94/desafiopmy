package main

import (
	"orderprocessor/internal/server"
)

func main() {
	sv := server.NewServer()
	sv.Run()
}
