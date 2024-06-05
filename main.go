package main

import (
	"fmt"
	s "kill-list/server"
)

func main() {
	fmt.Println("Fuck you!")
	server := s.NewServer(s.WithName("my-server"))
	fmt.Printf("%+v\n", server)
}
