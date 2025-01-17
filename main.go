package main

import (
	"anmit007/go-grpc/proto"
	"log"
)

func main() {
	person := proto.Person{
		Name: "anmit",
	}
	log.Println(person.GetName())
}
