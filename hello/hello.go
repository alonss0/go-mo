package main

import (
	"fmt"

	"alonss0/greetings"

	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(3)

	names := []string{"Connor", "Volkanoski", "Illa"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
