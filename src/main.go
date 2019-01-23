package main

import (
	"server"
	"client"
)

func main() {
	server.Start()
	client.Start()
}
