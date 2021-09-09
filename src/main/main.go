package main

import (
	"server"
)

func main() {
	server := server.CreateServer()
	server.Start()
}
