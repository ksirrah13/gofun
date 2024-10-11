package main

import (
	"fmt"
	"log"

	"github.com/ksirrah13/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("gofun: ")
	log.SetFlags(0)

	fmt.Printf("who is that? %s\n", quote.Go())
	greeting, err := greetings.GetGreeting("Kyle")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greeting)

	greetings.GreetMany("Kyle", "Steph", "Stan")
}
