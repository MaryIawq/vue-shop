package main

import (
	"log"
)

func main() {
	err := runSeed()
	if err != nil {
		log.Fatal(err)
	}
}

func runSeed() error {
	return nil
}
