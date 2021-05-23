package main

import (
	"github.com/goyder/bogglers/pkg/engine"
	"log"
	"os"
	"strconv"
)

func main() {

	// The dictionary path is the first argument. We read this in.
	dictionaryPath := os.Args[1]
	dictionaryData, err := os.Open(dictionaryPath)
	if err != nil {
		log.Fatal(err)
	}
	dictionary := engine.LoadDictionary(dictionaryData)
	dictionaryData.Close()

	// And we pull in configs
	minLetters, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	maxLetters, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	// Point of reflection: this very basic form of error checking seems incredibly simplistic and repetitive.
	// I bet there's a significant body of work in this space.

	// Spin up our inputs
	game := engine.BoggleGame{
		Network:       engine.GenerateNetwork(),
		LetterMapping: engine.GenerateRandomNetworkLetterMapping(1),
		Dictionary:    dictionary,
	}

	// And go
	words := engine.SolveBoggleNetwork(game, minLetters, maxLetters)

	// Outputs
	for _, word := range words {
		println(word)
	}
}
