package main

import (
	"fmt"
	s "kill-list/server"
	"log"
)

func main() {
	fmt.Println("Fuck you!")
	server := s.NewServer(s.WithName("my-server"))
	log.Fatalln(server.Listen(":8080"))
}
