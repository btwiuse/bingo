package main

import (
	"log"
	"os"

	"github.com/btwiuse/bingo"
)

func main() {
	err := bingo.Run(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}
