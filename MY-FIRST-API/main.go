package main

import (
	"MY-FIRST-API/internal/todo"
	"MY-FIRST-API/internal/transport"
	"log"
)

func main() {
	//Service Layer
	svc := todo.NewService()
	//Server Layer
	server := transport.NewServer(svc)
	//Serve
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
