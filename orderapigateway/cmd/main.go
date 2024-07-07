package main

import (
	"apigateway/internal/server"
)

func main() {
	sv := server.NewServer()
	sv.Run()
}
