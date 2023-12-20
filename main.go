package main

import (
	"log"
	"os"
	"wisdom/cmd"
)

func main() {
	log.Println("Main - application is starting....")
	serviceName := os.Getenv("SERVICE_NAME")
	if len(os.Args) > 1 {
		serviceName = os.Args[1]
	}

	switch serviceName {
	case "server":
		cmd.RunServer()
	case "client":
		cmd.RunClient()
	default:
		log.Println("Main - service name: ", serviceName)
		panic("invalid SERVICE_NAME ")
	}
}
