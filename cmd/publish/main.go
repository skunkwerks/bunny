package main

import (
	"bunny"
	"fmt"
	"log"
	"os"
)

func main() {
	b, err := bunny.NewBrokerFromArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.Publish())
}
